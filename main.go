package main

import (
	_ "net/http/pprof"
	"sso/engine"
	"sso/hooks"
	"sso/router"

	"github.com/gin-gonic/gin"
)

func main() {
	app := engine.New()

	runHTTP(app)

	app.WaitClose()
}

func runHTTP(app engine.Engine) {
	// mode
	gin.SetMode(gin.ReleaseMode)

	// disable
	gin.DisableConsoleColor()

	// init gin
	r := gin.Default()

	// hook
	r.Use(hooks.Session(), hooks.Auth())

	// router
	app.HTTP(router.HTTPConfig())

	// run monitor service
	go engine.RunHTTPMonitorService()

	// start http service
	engine.RunHTTPService(r)

}
