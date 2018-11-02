package xcontext

import (
	"net/http"
	"sso/xengine/xcode"
	"time"
)

// Context context
type Context interface {
	Request() *http.Request
	Response() *http.Response
	Error(int) Error
}

// XContext context
type XContext struct {
	startTime time.Time
	request   *http.Request
	response  *http.Response
}

// 实例化context
func newContextGin() Context {
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
	r := xcode.GetCode(code)

	//未定义的返回码直接返回
	if r == nil {
		return &XError{
			code: code,
		}
	}

	return &XError{
		code: r.Code,
		msg:  r.Msg,
		info: r.Info,
	}
}
