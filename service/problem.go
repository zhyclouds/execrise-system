package service

import (
	"execrise-system/define"
	"execrise-system/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags 公共方法
// @Summary 获取题目列表
// @param page query int false "page"
// @param size query int false "size"
// @param keyword query string false "keyword"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
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

	tx := models.GetProblemList(keyword)
	list := make([]*models.ProblemBasic, size)
	err = tx.Count(&count).Omit("content").Limit(size).Offset(page).Find(&list).Error
	if err != nil {
		log.Println("GetProblemList Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
