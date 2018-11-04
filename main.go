package main

import (
	_ "net/http/pprof"
	"sso/router"
	"sso/xengine"
	"sso/xengine/xdefine"
)

func main() {
	app := xengine.New()

	httpServer(app)

	app.Wait()
}

func httpServer(app xdefine.Engine) {
	config := router.HTTPWebConfig()
	app.ServerHTTPWeb(8001, config)
}

// func runHTTP(app engine.Engine) {
// 	// mode
// 	gin.SetMode(gin.ReleaseMode)

// 	// disable
// 	gin.DisableConsoleColor()

// 	// init gin
// 	r := gin.Default()

// 	// hook
// 	r.Use(hooks.Session(), hooks.Auth())

// 	// router
// 	app.HTTP(router.HTTPConfig())

// 	// run monitor service
// 	go engine.RunHTTPMonitorService()

// 	// start http service
// 	engine.RunHTTPService(r)

// }
