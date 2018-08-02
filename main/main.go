package main

import (
	"sso/config"
	"sso/hooks"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	router := gin.Default()

	router.Use(hooks.Auth())

	config.SetRouter(router)
	serverPort := config.GetServerPort()
	runHttp(serverPort, router)
}
