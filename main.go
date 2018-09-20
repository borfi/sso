package main

import (
	"log"
	_ "net/http/pprof"
	"sso/engine"
	"sso/engine/xconfig"
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

	// get monitor port
	monitorPort, err := xconfig.Config().Int("service", "monitor_port")
	if err != nil {
		log.Fatalf("get service monitor port err: %v", err)
		return
	}

	// run monitor service
	go engine.HTTPMonitorService(monitorPort)

	// get server port
	servicePort, err := xconfig.Config().Int("service", "port")
	if err != nil {
		log.Fatalf("get service port err: %v", err)
		return
	}

	// start http service
	engine.HTTPService(r, servicePort)
}
