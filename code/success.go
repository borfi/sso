package code

import "sso/engine"

const (
	// Success .
	Success = 1
)

func init() {
	engine.RegisterCode([]engine.CodeItem{
		{Code: Success, Msg: "成功"},
	})
}
