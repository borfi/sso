package engine

import "fmt"

// XError .
type XError struct {
	code int
	msg  string
	info string
}

// Format .
func (e *XError) Format(fields ...interface{}) Error {
	if len(fields) > 0 {
		e.msg = fmt.Sprintf(e.msg, fields)
		e.info = fmt.Sprintf(e.info, fields)
		return e
	}
	return e
}

// Code .
func (e *XError) Code() int {
	return e.code
}

// Msg .
func (e *XError) Msg() string {
	return e.msg
}

// Info .
func (e *XError) Info() string {
	return e.info
}
