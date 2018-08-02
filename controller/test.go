package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Test ...
func Test(c *gin.Context) {
	log.Println(c.Query("name"))
	c.JSON(200, gin.H{
		"message": "controller test ..",
	})
}

//Test2 ...
func Test2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "controller test2 ..",
	})
}
