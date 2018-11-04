package router

import (
	"sso/controller"
	"sso/hooks"
	"sso/xengine/xdefine"

	"github.com/gin-gonic/gin"
)

//HTTPWebConfig HTTP勾子和路由配置
func HTTPWebConfig() func(xdefine.Server, *gin.Engine) {
	return func(s xdefine.Server, r *gin.Engine) {
		// hook
		r.Use(hooks.Session(), hooks.Auth())

		//系统
		// rsystem := r.Group("/system")
		// {
		// 	rsystem.GET("/analysis-config", controller.AnalysisConfig)
		// }

		//测试
		rtest := r.Group("/test")
		{
			rtest.GET("/test", s.Handler(controller.Test).(gin.HandlerFunc))
			//rtest.GET("/session-set", controller.TestSessionSet)
			//rtest.GET("/session-get", controller.TestSessionGet)
		}
	}
}
