package hooks

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Init 初始化
func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		executeTime(c)
		session(c)
		c.Next()
	}

}

func executeTime(c *gin.Context) {
	c.Set("sso_exec_start_time", time.Now())
}

func session(c *gin.Context) {

}
