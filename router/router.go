package router

import (
	"sso/controller"

	"github.com/gin-gonic/gin"
)

//Set 配置路由
func Set(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/test", controller.Test)
		v1.GET("/test2", controller.Test2)
	}
}
