package engine

import "fmt"

// Format 格式化错误信息
func (e *xError) Format(fields ...interface{}) Error {
	if len(fields) > 0 {
		e.msg = fmt.Sprintf(e.msg, fields)
		e.info = fmt.Sprintf(e.info, fields)
		return e
	}
	return e
}

// SetError 设置具体的错误信息
func (e *xError) SetError(err error) Error {
	e.err = err
	return e
}

// Code 取错误码
func (e *xError) Code() int {
	return e.code
}

// Msg 取返回给外界的错误信息
func (e *xError) Msg() string {
	return e.msg
}

// Info 取返回开发人员的错误信息
func (e *xError) Info() string {
	return e.info
}

// Error 取具体的错误信息
func (e *xError) Error() error {
	return e.err
}
