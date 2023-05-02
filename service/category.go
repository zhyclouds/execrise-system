package service

import (
	"execrise-system/define"
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
