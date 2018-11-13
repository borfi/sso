package httpapi

import (
	"fmt"
	"sso/engine/xrecovery"
	"sso/engine/xservice"
)

// Recovery 拦截panic
func Recovery() xservice.Handler {
	return func(c xservice.Context) (interface{}, xservice.Error) {
		defer func() {
			if err := recover(); err != nil {
				stack := xrecovery.Stack(3)
				fmt.Printf("[Recovery] %s\n", string(stack))
				return
			}
		}()
		//fmt.Println("this is Recovery hook")
		return nil, nil
	}
}
