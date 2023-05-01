package router

import (
	_ "execrise-system/docs"
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

	return r
}
