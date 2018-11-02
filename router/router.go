package router

import (
	"sso/controller"
	"sso/engine"

	"github.com/gin-gonic/gin"
)

//HTTPConfig HTTP路由配置
func HTTPConfig() func(*gin.Engine) {
	return func(r *gin.Engine) {
		//系统
		// rsystem := r.Group("/system")
		// {
		// 	rsystem.GET("/analysis-config", controller.AnalysisConfig)
		// }

		//测试
		rtest := r.Group("/test")
		{
			rtest.GET("/test", engine.Get().HandlerGin(controller.Test))
			//rtest.GET("/session-set", controller.TestSessionSet)
			//rtest.GET("/session-get", controller.TestSessionGet)
		}
	}
}
