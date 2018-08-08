package controller

import (
	"log"
	"sso/engine/xconfig"

	"github.com/gin-gonic/gin"
)

//Test ...
func Test(c *gin.Context) {
	log.Println(c.Query("name"))

	port, err := xconfig.Config().String("service", "port")

	c.JSON(200, gin.H{
		"message": "controller test ..",
		"port":    port,
		"err":     err,
	})
}

//Test2 ...
func Test2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "controller test2 ..",
	})
}
