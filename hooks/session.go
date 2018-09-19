package hooks

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Session session
func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("这里进行session初始化")
		c.Next()
	}
}
