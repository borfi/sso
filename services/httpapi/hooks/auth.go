package httpapi

import (
	"sso/engine/xservice"
)

//Auth 认证
func Auth() xservice.Handler {
	return func(xservice.Context) (interface{}, xservice.Error) {
		//fmt.Println("this is Auth hook")
		return nil, nil
	}
}
