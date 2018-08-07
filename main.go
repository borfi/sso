package main

import (
	"sso/hooks"
	"sso/router"
	"sso/xapp/xrun"

	"github.com/gin-gonic/gin"
)

func main() {
	// mode
	gin.SetMode(gin.ReleaseMode)

	// disable
	gin.DisableConsoleColor()

	// init
	r := gin.Default()

	// hook
	r.Use(hooks.Auth())

	// router
	router.Set(r)

	// service
	xrun.HTTP(r)
}
