package xerror

import "fmt"

// Error 错误接口
type Error interface {
	Code() int
	Msg() string
	Info() string
	SetError(error) Error
	Format(...interface{}) Error
}

// XError 错误
type XError struct {
	code int
	msg  string
	info string
	err  error
}

// SetError 设置具体的错误信息
func (e *XError) SetError(err error) Error {
	e.err = err
	return e
}

// Code 取错误码
func (e *XError) Code() int {
	return e.code
}

// Msg 取返回给外界的错误信息
func (e *XError) Msg() string {
	return e.msg
}

// Info 取返回开发人员的错误信息
func (e *XError) Info() string {
	return e.info
}

// Error 取具体的错误信息
func (e *XError) Error() error {
	return e.err
}

// Format 格式化错误信息
func (e *XError) Format(fields ...interface{}) Error {
	if len(fields) > 0 {
		e.msg = fmt.Sprintf(e.msg, fields)
		e.info = fmt.Sprintf(e.info, fields)
		return e
	}
	return e
}
