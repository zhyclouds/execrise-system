package service

import (
	"execrise-system/define"
	"execrise-system/helper"
	"execrise-system/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetCategoryList
// @Tags 管理员私有方法
// @Summary 获取分类列表
// @param authorization header string true "authorization"
// @param page query int false "page"
// @param size query int false "size"
// @param keyword query string false "keyword"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /category-list [get]
func GetCategoryList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page Strconv Error:", err)
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetProblemList Size Strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64

	keyword := c.Query("keyword")

	categoryList := make([]*models.CategoryBasic, size)
	err = models.DB.Model(&models.CategoryBasic{}).Where("name like ?", "%"+keyword+"%").
		Count(&count).Limit(size).Offset(page).Find(&categoryList).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "获取分类列表失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"count": count,
			"list":  categoryList,
		},
	})
}

// AddCategory
// @Tags 管理员私有方法
// @Summary 分类创建
// @param authorization header string true "authorization"
// @param name formData string true "name"
// @param parentId formData string false "parentId"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /add-category [post]
func AddCategory(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))

	category := &models.CategoryBasic{
		Identity: helper.GetUUid(),
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Create(&category).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "创建分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "创建分类成功",
	})
}

// UpdateCategory
// @Tags 管理员私有方法
// @Summary 分类修改
// @param authorization header string true "authorization"
// @param identity header string true "identity"
// @param name formData string true "name"
// @param parentId formData string false "parentId"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /update-category [put]
func UpdateCategory(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	identity := c.PostForm("identity")
	if name == "" || identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "参数错误",
		})
		return
	}

	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Model(&models.CategoryBasic{}).Where("identity = ?", identity).Updates(&category).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "修改分类失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "修改分类成功",
	})
}

// DelCategory
// @Tags 管理员私有方法
// @Summary 分类删除
// @param authorization header string true "authorization"
// @param identity query string true "identity"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /del-category [delete]
func DelCategory(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "参数错误",
		})
		return
	}

	var cnt int64
	err := models.DB.Model(&models.ProblemCategory{}).Where("category_id = (SELECT id FROM category_basic WHERE identity = ? ) ", identity).Count(&cnt).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "获取分类关联的问题失败",
		})
		return
	}

	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "该分类下存在问题，无法删除",
		})
		return
	}

	err = models.DB.Where("identity = ?", identity).Delete(&models.CategoryBasic{}).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "删除分类失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "删除分类成功",
	})
}
