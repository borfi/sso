package xservice

import (
	"fmt"
	"sso/engine/xcode"
)

// xError 错误
type xError struct {
	code   int                    // 错误码
	msg    string                 // 给外界看的信息
	info   string                 // 给内部看的信息
	err    error                  // 具体错误
	fields map[string]interface{} // 相关数据
}

// 创建一个错误
func newError(code int, err error) Error {
	r := xcode.Get(code)

	//未定义的返回码直接返回
	if r == nil {
		return &xError{
			code:   code,
			err:    err,
			fields: nil,
		}
	}

	return &xError{
		code:   r.Code,
		msg:    r.Msg,
		info:   r.Info,
		err:    err,
		fields: nil,
	}
}

// GetCode 取错误码
func (e *xError) GetCode() int {
	return e.code
}

// GetMsg 取给外界看的错误信息
func (e *xError) GetMsg() string {
	return e.msg
}

// GetInfo 取给内部看的错误信息
func (e *xError) GetInfo() string {
	return e.info
}

// GetError 取具体的错误
func (e *xError) GetError() error {
	return e.err
}

// GetFields 取相关的数据
func (e *xError) GetFields() map[string]interface{} {
	return e.fields
}

// SettFields 设置相关数据
func (e *xError) SetFields(fs map[string]interface{}) Error {
	if e.fields == nil {
		e.fields = make(map[string]interface{}, 0)
	}
	for k := range fs {
		e.fields[k] = fs[k]
	}
	return e
}

// Format 格式化错误信息（内部+外界）
func (e *xError) Format(fields []interface{}) Error {
	if len(fields) == 0 {
		return e
	}
	e.msg = fmt.Sprintf(e.msg, fields...)
	e.info = fmt.Sprintf(e.info, fields...)
	return e
}
