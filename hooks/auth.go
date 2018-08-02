package hooks

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Auth 权限认证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("aaaaaaaa")
		c.Next()
		log.Println("bbbbbbbb")
	}
}
