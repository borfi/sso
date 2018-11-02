package xerror

import (
	"fmt"
	"sso/xengine/xcode"
	"sso/xengine/xdefine"
)

// XError 错误
type XError struct {
	code int    //错误码
	msg  string //可以给外界看的信息
	info string //可以给内部看的信息
	err  error  //具体错误
}

// New 创建一个错误
func New(code int) xdefine.Error {
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

// Code 取错误码
func (e *XError) Code() int {
	return e.code
}

// Msg 取可以给外界看的错误信息
func (e *XError) Msg() string {
	return e.msg
}

// Info 取可以给内部看的错误信息
func (e *XError) Info() string {
	return e.info
}

// Error 取具体的错误
func (e *XError) Error() error {
	return e.err
}

// SetError 设置具体的错误
func (e *XError) SetError(err error) xdefine.Error {
	e.err = err
	return e
}

// Format 格式化错误信息（内部+外界）
func (e *XError) Format(fields ...interface{}) xdefine.Error {
	if len(fields) > 0 {
		e.msg = fmt.Sprintf(e.msg, fields)
		e.info = fmt.Sprintf(e.info, fields)
		return e
	}
	return e
}
