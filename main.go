package main

import (
	"sso/hooks"
	"sso/router"
	"sso/xapp/xconfig"
	"sso/xapp/xrun"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(hooks.Auth())

	router.Set(r)
	serverPort := xconfig.Xconf().ServerPort()
	xrun.HTTP(serverPort, r)
}
