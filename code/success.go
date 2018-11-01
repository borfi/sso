package code

import "sso/engine"

const (
	// Success .
	Success = 1
)

func init() {
	engine.CodeRegister([]engine.XCode{
		{Code: Success, Msg: "成功"},
	})
}
