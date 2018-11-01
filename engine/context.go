package engine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// XContext .
type XContext struct {
	gin.Context
}

// 实例化context
func newContext() Context {
	return &XContext{}
}

// Test .
func (ctx *XContext) Test() {
	fmt.Println("this is mycontext test func")
}

// Request 请求
func (ctx *XContext) Request() *http.Request {

	return nil
}

// Response 响应
func (ctx *XContext) Response() *http.Response {

	return nil
}

// Error .
func (ctx *XContext) Error(code int) Error {
	r := engine.codes.Get(code)

	//未定义的返回码直接返回
	if r == nil {
		return &XError{
			code: code,
			msg:  "",
			info: "",
		}
	}

	return &XError{
		code: r.Code,
		msg:  r.Msg,
		info: r.Info,
	}
}
