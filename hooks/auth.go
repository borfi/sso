package hooks

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Auth 认证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("这里做认证")
	}
}
