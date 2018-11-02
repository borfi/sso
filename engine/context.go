package engine

import (
	"net/http"
)

// 实例化context
func newContext() Context {
	return &XContext{}
}

// Request 请求
func (ctx *XContext) Request() *http.Request {
	return ctx.request
}

// Response 响应
func (ctx *XContext) Response() *http.Response {
	return ctx.response
}

// Error 组装错误返回
func (ctx *XContext) Error(code int) Error {
	r := GetCode(code)

	//未定义的返回码直接返回
	if r == nil {
		return &xError{
			code: code,
		}
	}

	return &xError{
		code: r.Code,
		msg:  r.Msg,
		info: r.Info,
	}
}
