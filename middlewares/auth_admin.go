package middlewares

import (
	"execrise-system/helper"
	"github.com/gin-gonic/gin"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaims, err := helper.AnalyseToken(auth)
		if err != nil {
			c.JSON(200, gin.H{
				"code": -1,
				"data": "invalid authorization",
			})
			c.Abort()
			return
		}
		if userClaims.IsAdmin != 1 {
			c.JSON(200, gin.H{
				"code": -1,
				"data": "authorization not admin",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
