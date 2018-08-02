package config

import (
	"sso/controller"

	"github.com/gin-gonic/gin"
)

//SetRouter 配置路由
func SetRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		t := router.Group("/v1/t")
		{
			t.GET("/t1", controller.Test)
			t.GET("/t2", controller.Test2)
		}
		v1.GET("/test", controller.Test)
		v1.GET("/test2", controller.Test2)
	}
}
