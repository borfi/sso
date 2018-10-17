package main

import (
	_ "net/http/pprof"
	"sso/engine"
	"sso/hooks"
	"sso/router"

	"github.com/gin-gonic/gin"
)

func main() {
	runHTTP()
}

func runHTTP() {
	// mode
	gin.SetMode(gin.ReleaseMode)

	// disable
	gin.DisableConsoleColor()

	// init gin
	r := gin.Default()

	// hook
	r.Use(hooks.Session(), hooks.Auth())

	// router
	router.Set(r)

	// run monitor service
	go engine.RunHTTPMonitorService()

	// start http service
	engine.RunHTTPService(r)
}
