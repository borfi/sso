package main

import (
	"net/http"
	_ "net/http/pprof"
	"sso/engine"
	"sso/hooks"
	"sso/router"
	"sso/xengine"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := xengine.New()

	httpServer(app)

	app.WaitClose()
}

func httpServer(app xengine.Engine) {
	router := router.HTTPConfig()
	app.ServerHTTP(&http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})
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
