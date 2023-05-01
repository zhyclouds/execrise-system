package service

import (
	"execrise-system/helper"
	"execrise-system/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// GetUserDetail
// @Tags 公共方法
// @Summary 用户详情
// @param identity query string false "user identity"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "identity is empty",
		})
		return
	}

	data := new(models.UserBasic)
	err := models.DB.Omit("password").Where("identity = ?", identity).First(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// Login
// @Tags 公共方法
// @Summary 用户登录
// @param username formData string false "username"
// @param password formData string false "password"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "username or password is empty",
		})
		return
	}

	// md5加密
	password = helper.GetMd5(password)

	data := new(models.UserBasic)
	err := models.DB.Where("name = ? AND password = ? ", username, password).First(data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}

	token, err := helper.GenerateToken(data.Identity, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @param email formData string false "email"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "email is empty",
		})
		return
	}
	code := helper.GetRand()
	models.RDB.Set(c, email, code, time.Second*300)
	err := helper.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "验证码发送成功",
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @param name formData string true "name"
// @param password formData string true "password"
// @param phone formData string false "phone"
// @param email formData string true "email"
// @param code formData string true "code"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	code := c.PostForm("code")

	if email == "" || password == "" || name == "" || code == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "参数不正确",
		})
		return
	}

	// 验证码校验
	sysCode, err := models.RDB.Get(c, email).Result()
	if err != nil {
		log.Printf("Get code error: %v \n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "请重新获取验证码",
		})
		return
	}
	if sysCode != code {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "验证码不正确",
		})
		return
	}
	// 判断用户是否已经存在
	var count int64
	err = models.DB.Model(&models.UserBasic{}).Where("mail = ?", email).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "该邮箱已被注册",
		})
		return
	}

	// md5加密
	data := models.UserBasic{
		Identity: helper.GetUUid(),
		Name:     name,
		Password: helper.GetMd5(password),
		Phone:    phone,
		Mail:     email,
	}

	// 数据库操作
	err = models.DB.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}

	// 生成token
	token, err := helper.GenerateToken(data.Identity, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "Generate Token Error:" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"msg":   "注册成功",
			"token": token,
		},
	})
}
