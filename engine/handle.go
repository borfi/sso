package engine

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// HandlerGin .
func (x *XEngine) HandlerGin(h func(Context) (interface{}, Error)) gin.HandlerFunc {
	return func(c *gin.Context) {
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
