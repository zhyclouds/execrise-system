package router

import (
	_ "execrise-system/docs"
	"execrise-system/middlewares"
	"execrise-system/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// TODO: 配置路由规则
	// 问题
	r.GET("/problem-list", service.GetProblemList)

	// 用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)
	r.GET("/rank-list", service.GetRankList)

	// 提交
	r.GET("/submit-list", service.GetSubmitList)

	// 管理员私有方法
	r.POST("/add-problem", middlewares.AuthAdminCheck(), service.AddProblem)

	// 分类列表
	r.GET("/category-list", middlewares.AuthAdminCheck(), service.GetCategoryList)

	// 分类创建、删除、修改
	// 分类创建
	r.POST("/add-category", middlewares.AuthAdminCheck(), service.AddCategory)
	// 分类删除
	r.DELETE("del-category", middlewares.AuthAdminCheck(), service.DelCategory)
	// 分类修改
	r.PUT("/update-category", middlewares.AuthAdminCheck(), service.UpdateCategory)

	return r
}
