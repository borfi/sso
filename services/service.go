package services

import (
	"fmt"
	"sso/engine/xservice"
	httpapi "sso/services/httpapi/hooks"
	httpweb "sso/services/httpweb/hooks"

	"github.com/gin-gonic/gin"
)

// Register 注册服务函数
func Register() func() {
	return func() {
		fmt.Println("web init")
		xservice.RegisterWeb(&xservice.WebConfig{
			Name:                 "测试web服务",                                                                // 服务名称
			Port:                 8001,                                                                     // 监听端口
			ReadTimeout:          1,                                                                        // 读超时时间
			WriteTimeout:         2,                                                                        // 写超时时间
			MaxHeaderBytes:       1 << 20,                                                                  // 请求的头域最大允许长度 1M
			GracefullQuitTimeout: 5,                                                                        // 优雅退出的超时时间
			Hooks:                []gin.HandlerFunc{httpweb.Recovery(), httpweb.Session(), httpweb.Auth()}, // 钩子
			Router:               httpWebRouter(),                                                          // 路由
			TemplatePath:         "views/**/*",                                                             // 模板文件路径
		})

		fmt.Println("api init")
		xservice.RegisterHTTP(&xservice.HTTPConfig{
			Name:                 "测试api服务",                                              // 服务名称
			Port:                 8002,                                                   // 端口
			ReadTimeout:          1,                                                      // 读超时时间
			WriteTimeout:         2,                                                      // 写超时时间
			MaxRequestBodySize:   1 << 20,                                                // 请求的头域最大允许长度 1M
			GracefullQuitTimeout: 5,                                                      // 优雅退出的超时时间
			Hooks:                []xservice.Handler{httpapi.Recovery(), httpapi.Auth()}, // 钩子
			Router:               httpAPIRouter(),                                        // 路由
		})
	}
}
