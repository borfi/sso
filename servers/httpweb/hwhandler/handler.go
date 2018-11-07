package hwhandler

import (
	"github.com/gin-gonic/gin"
)

// Handler web处理器
func Handler(f func(*gin.Context)) gin.HandlerFunc {
	return func(g *gin.Context) {
		f(g)
	}
}
