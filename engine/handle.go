package engine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerGin .
func (x *xEngine) HandlerGin(f func(Context) (interface{}, Error)) gin.HandlerFunc {
	return func(gctx *gin.Context) {
		ctx := x.CtxGet()
		defer x.CtxPut(ctx)

		data, xerr := f(ctx)
		if xerr != nil {
			fmt.Println("err:", xerr)
		} else {
			fmt.Println("succ:", data)
		}

		gctx.JSON(http.StatusOK, responseJSON(ctx, data, xerr))
		return
	}
}
