package services

import (
	httpapi "sso/services/httpapi/controllers"
	httpweb "sso/services/httpweb/controllers"

	"github.com/gin-gonic/gin"
)

// httpAPIRouter HTTP api 路由配置
func httpAPIRouter() func(*gin.Engine) {
	return func(r *gin.Engine) {
		//测试
		rtest := r.Group("/test")
		{
			rtest.GET("/test", httpAPIHandler(httpapi.Test))
		}
	}

}

// httpWebRouter HTTP web 路由配置
func httpWebRouter() func(*gin.Engine) {
	return func(r *gin.Engine) {
		//测试
		rtest := r.Group("/test")
		{
			rtest.GET("/test", httpWebHandler(httpweb.Test))
			//rtest.GET("/session-set", controller.TestSessionSet)
			//rtest.GET("/session-get", controller.TestSessionGet)
		}
	}

}
