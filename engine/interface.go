package engine

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
}

// Context context
type Context interface {
	Request() *http.Request
	Response() *http.Response
	Error(int) Error
}

// Error 错误
type Error interface {
	Code() int
	Msg() string
	Info() string
	SetError(error) Error
	Format(...interface{}) Error
}
