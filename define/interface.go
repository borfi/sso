package define

import "sso/xengine/xerror"

// Error 错误接口
type Error interface {
	Code() int            //返回错误码
	Msg() string          //返回可以给外界看的信息
	Info() string         //返回可以给内部看的信息
	Error() error         //取错误
	SetError(error) Error //设置错误
	SetFields(*xerror.Fields) *Error
	Format(...interface{}) Error //格式化错误信息（内部+外界）
}
