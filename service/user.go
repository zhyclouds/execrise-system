package service

import (
	"execrise-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

	// TODO: 从数据库中获取用户信息
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
