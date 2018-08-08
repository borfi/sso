package code

import "sso/engine"

const (
	Success = 100000
)

func init() {
	engine.RegisterCode([]engine.CodeItem{
		{Code: Success, Msg: "成功"},
	})
}
