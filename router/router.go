package router

import (
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
		rtest.GET("/test2", controller.Test2)
	}
}
