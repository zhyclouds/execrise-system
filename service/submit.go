package service

import (
	"execrise-system/define"
	"execrise-system/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 提交列表
// @param page query int false "page"
// @param size query int false "size"
// @param problem_identity query string false "problem_identity"
// @param user_identity query string false "user_identity"
// @param status query string false "status"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetSubmitList Size Strconv Error:", err)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetSubmitList Page Strconv Error:", err)
		return
	}
	page = (page - 1) * size

	var count int64
	list := make([]*models.SubmitBasic, size)

	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.Query("status"))
	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	err = tx.Count(&count).Omit("content").Limit(size).Offset(page).Find(&list).Error
	if err != nil {
		log.Println("GetSubmitList Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"count": count,
			"list":  list,
		},
	})
}
