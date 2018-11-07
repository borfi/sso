package main

import (
	_ "net/http/pprof"
	"sso/servers/httpapi"
	"sso/servers/httpweb"
	"time"
)

func main() {
	//app := xengine.New()

	httpweb.Run()

	httpapi.Run()

	time.Sleep(100 * time.Second)

	//app.Wait()
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
