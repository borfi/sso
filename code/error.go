package code

import "sso/engine/xcode"

const (
	// ParamsError .
	ParamsError = 500100
)

func init() {
	xcode.Register([]xcode.Code{
		{Code: ParamsError, Msg: "参数%v有误", Info: "params error"},
	})
}
