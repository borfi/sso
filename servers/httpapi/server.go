package httpapi

import (
	"sso/servers/httpapi/hooks"
	"sso/servers/httpapi/router"
	"sso/xengine/xserver"
)

// Run 运行服务
func Run() {
	app := xserver.NewHTTP()
	app.Hooks(hooks.Recovery(), hooks.Auth())
	app.Router(router.Config())
	app.Run(xserver.Config{
		ServerName:           "测试接口服务", //服务名
		ServerPort:           8002,     //端口
		ReadTimeout:          1,        // 读超时时间
		WriteTimeout:         2,        // 写超时时间
		MaxHeaderBytes:       1 << 20,  // 请求的头域最大允许长度 1M
		GracefullExitTimeout: 5,        // 优雅关闭的超时时间
	})
}
