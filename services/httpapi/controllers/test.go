package httpapi

import (
	"github.com/gin-gonic/gin"
)

//Test ...
func Test(c *gin.Context) {
	//port, _ := xconfig.Config().String("service", "port")

	c.JSON(200, gin.H{
		"code":    1,
		"message": "this is a test",
		"info":    "hahaha",
	})
}
