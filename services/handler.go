package service

import (
	"github.com/gin-gonic/gin"
)

// httpAPIHandler http api 处理器
func httpAPIHandler(f func(*gin.Context)) gin.HandlerFunc {
	return func(g *gin.Context) {
		f(g)
	}
}

// httpWebHandler http web 处理器
func httpWebHandler(f func(*gin.Context)) gin.HandlerFunc {
	return func(g *gin.Context) {
		f(g)
	}
}
