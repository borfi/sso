package main

import (
	"sso/config"
	"sso/hooks"
	"sso/run"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	router := gin.Default()

	router.Use(hooks.Auth())

	config.SetRouter(router)
	serverPort := config.GetServerPort()
	run.Http(serverPort, router)
}
