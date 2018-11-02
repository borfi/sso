package xengine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Engine 引擎控制器
type Engine interface {
	Status() bool
	WaitClose()
	Close(*XContext)
	CtxGet() *XContext
	CtxPut(*XContext)
	HandlerGin(func(Context) (interface{}, Error)) gin.HandlerFunc
	HTTP(router *gin.Engine) *http.Server
}
