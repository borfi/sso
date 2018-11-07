package httpweb

import (
	"sso/servers/httpweb/hwhooks"
	"sso/servers/httpweb/hwrouter"
	"sso/xengine/xserver"
)

// Run 运行服务
func Run() {
	app := xserver.NewHTTP()
	app.Hooks(hwhooks.Recovery(), hwhooks.Session(), hwhooks.Auth())
	app.Router(hwrouter.Config())
	app.LoadHTMLGlob("views/**/*")
	app.Run(xserver.Config{
		ServerName:           "测试web服务", //服务名
		ServerPort:           8001,      //端口
		ReadTimeout:          1,         // 读超时时间
		WriteTimeout:         2,         // 写超时时间
		MaxHeaderBytes:       1 << 20,   // 请求的头域最大允许长度 1M
		GracefullExitTimeout: 5,         // 优雅关闭的超时时间
	})
}
