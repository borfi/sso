package main

import (
	"sso/engine"
	"sso/hooks"
	"sso/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// mode
	gin.SetMode(gin.ReleaseMode)

	// disable
	gin.DisableConsoleColor()

	// init gin
	r := gin.Default()

	// hook
	r.Use(hooks.Auth())

	// router
	router.Set(r)

	// init engin
	app := engine.New()

	// service
	app.HTTP(r)
}
