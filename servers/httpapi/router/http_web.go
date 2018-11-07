package router

import (
	"sso/servers/httpapi/controllers"
	"sso/servers/httpapi/handler"

	"github.com/gin-gonic/gin"
)

// Config HTTP路由配置
func Config() func(*gin.Engine) {
	return func(r *gin.Engine) {
		//测试
		rtest := r.Group("/test")
		{
			rtest.GET("/test", handler.Handler(controllers.Test))
		}
	}

}
