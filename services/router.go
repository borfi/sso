package services

import (
	"sso/engine/xservice"
	api "sso/services/api/controllers"
	web "sso/services/web/controllers"

	"github.com/gin-gonic/gin"
)

// httpAPIRouter HTTP api 路由配置
func httpAPIRouter() []*xservice.Router {
	r := []*xservice.Router{
		{
			Method:  "GET",
			Path:    "/test/test",
			Handler: api.Test,
		},
	}
	return r
}

// webRouter web 路由配置
func webRouter() func(*gin.Engine) {
	return func(r *gin.Engine) {
		//测试
		rtest := r.Group("/test")
		{
			rtest.GET("/test", web.Test)
			//rtest.GET("/session-set", web.TestSessionSet)
			//rtest.GET("/session-get", web.TestSessionGet)
		}
	}
}
