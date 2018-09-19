package router

import (
	"sso/controller"

	"github.com/gin-gonic/gin"
)

//Set 配置路由
func Set(r *gin.Engine) {
	//系统
	rsystem := r.Group("/system")
	{
		rsystem.GET("/analysis-config", controller.AnalysisConfig)
	}

	//测试
	rtest := r.Group("/test")
	{
		rtest.GET("/test", controller.Test)
		rtest.GET("/session-set", controller.TestSessionSet)
		rtest.GET("/session-get", controller.TestSessionGet)
	}
}
