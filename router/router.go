package router

import (
	"sso/controller"

	"github.com/gin-gonic/gin"
)

//Set 配置路由
func Set(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/test", controller.Test)
		v1.GET("/test2", controller.Test2)
	}
}
