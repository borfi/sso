package engine

import (
	"fmt"
)

// HandlerFunc 返回接口函数类型
type HandlerFunc func(*XContext)

// Handler .
func (x *XEngine) Handler(h func(Context) (interface{}, Error)) HandlerFunc {
	return func(c *XContext) {
		ctx := x.CtxGet()
		data, err := h(ctx)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println("succ:", data)
		}

		x.CtxPut(ctx)
		return
	}
}
