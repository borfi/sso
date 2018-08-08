package router

import (
	"sso/controller/systemcontroller"
	"sso/controller/testcontroller"

	"github.com/gin-gonic/gin"
)

//Set 配置路由
func Set(r *gin.Engine) {
	//系统
	rsystem := r.Group("/system")
	{
		rsystem.GET("/analysis-config", systemcontroller.AnalysisConfig)
	}

	//测试
	rtest := r.Group("/test")
	{
		rtest.GET("/test", testcontroller.Test)
		rtest.GET("/test2", testcontroller.Test2)
	}
}
