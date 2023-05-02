package service

import (
	"encoding/json"
	"execrise-system/define"
	"execrise-system/helper"
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

// AddProblem
// @Tags 管理员私有方法
// @Summary 创建问题
// @param authorization header string true "authorization"
// @param title formData string true "title"
// @param content formData string true "content"
// @param max_runtime formData string true "max_runtime"
// @param max_mem formData string true "max_mem"
// @param category_ids formData string true "category_ids"
// @param test_cases formData string true "test_cases"
// @Success 200 {string} json "{"code":200,"data":""}"
// @Router /add-problem [post]
func AddProblem(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMem, _ := strconv.Atoi(c.PostForm("max_mem"))
	categoryIds := c.PostFormArray("category_ids")
	testCases := c.PostFormArray("test_cases")

	if title == "" || content == "" || len(testCases) == 0 || len(categoryIds) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "参数不能为空",
		})
		return
	}

	identity := helper.GetUUid()
	data := models.ProblemBasic{
		Identity:   identity,
		Title:      title,
		Content:    content,
		MaxRuntime: maxRuntime,
		MaxMem:     maxMem,
	}

	categoryBasics := make([]*models.ProblemCategory, 0)
	for _, id := range categoryIds {
		categoryId, _ := strconv.Atoi(id)
		categoryBasics = append(categoryBasics, &models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(categoryId),
		})
	}
	// TODO: 处理分类

	// 处理测试用例
	testCaseBasics := make([]*models.TestCase, 0)
	for _, testCase := range testCases {
		caseMap := make(map[string]string)
		err := json.Unmarshal([]byte(testCase), &caseMap)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": "测试用例格式错误",
			})
			return
		}
		if _, ok := caseMap["input"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": "测试用例输入不能为空",
			})
			return
		}
		if _, ok := caseMap["output"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": "测试用例输出不能为空",
			})
			return
		}
		testCaseBasics = append(testCaseBasics, &models.TestCase{
			Identity:        helper.GetUUid(),
			ProblemIdentity: identity,
			Input:           caseMap["input"],
			Output:          caseMap["output"],
		})
	}
	data.TestCases = testCaseBasics

	// 创建问题
	err := models.DB.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "创建问题失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}
