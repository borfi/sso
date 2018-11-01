package engine

import "net/http"

// Context context接口
type Context interface {
	Request() *http.Request
	Response() *http.Response
	Error(code int) Error
}

// Error 错误接口
type Error interface {
	Code() int
	Msg() string
	Info() string
	Format(fields ...interface{}) Error
}
