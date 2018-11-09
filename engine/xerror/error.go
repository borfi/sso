package xerror

import (
	"fmt"
	"sso/engine/xcode"
)

// Error 错误
type Error struct {
	code   int                    // 错误码
	msg    string                 // 给外界看的信息
	info   string                 // 给内部看的信息
	err    error                  // 具体错误
	fields map[string]interface{} // 相关数据
}

// New 创建一个错误
func New(code int) *Error {
	r := xcode.Get(code)

	//未定义的返回码直接返回
	if r == nil {
		return &Error{
			code:   code,
			err:    nil,
			fields: nil,
		}
	}

	return &Error{
		code:   r.Code,
		msg:    r.Msg,
		info:   r.Info,
		err:    nil,
		fields: nil,
	}
}

// Code 取错误码
func (e *Error) Code() int {
	return e.code
}

// Msg 取给外界看的错误信息
func (e *Error) Msg() string {
	return e.msg
}

// Info 取给内部看的错误信息
func (e *Error) Info() string {
	return e.info
}

// Error 取具体的错误
func (e *Error) Error() error {
	return e.err
}

// Fields 取具体的数据
func (e *Error) Fields() map[string]interface{} {
	return e.fields
}

// SetError 设置具体的错误
func (e *Error) SetError(err error) *Error {
	e.err = err
	return e
}

// PutFields 设置相关数据
func (e *Error) PutFields(fs map[string]interface{}) *Error {
	if e.fields == nil {
		e.fields = make(map[string]interface{}, 0)
	}
	for k := range fs {
		e.fields[k] = fs[k]
	}
	return e
}

// Format 格式化错误信息（内部+外界）
func (e *Error) Format(fields ...interface{}) *Error {
	if len(fields) == 0 {
		return e
	}
	e.msg = fmt.Sprintf(e.msg, fields)
	e.info = fmt.Sprintf(e.info, fields)
	return e
}
