package hwrouter

import (
	"sso/servers/httpweb/hwcontrollers"
	"sso/servers/httpweb/hwhandler"

	"github.com/gin-gonic/gin"
)

// Config HTTP路由配置
func Config() func(*gin.Engine) {
	return func(r *gin.Engine) {
		//测试
		rtest := r.Group("/test")
		{
			rtest.GET("/test", hwhandler.Handler(hwcontrollers.Test))
			//rtest.GET("/session-set", controller.TestSessionSet)
			//rtest.GET("/session-get", controller.TestSessionGet)
		}
	}

}
