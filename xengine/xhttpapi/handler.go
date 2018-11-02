package xhttpweb

import (
	"fmt"
	"net/http"
	"sso/xengine/xdefine"
	"sso/xengine/xresponse"

	"github.com/gin-gonic/gin"
)

// Handler api处理器
func (xh *XHTTP) Handler(f func(xdefine.Context) (interface{}, xdefine.Error)) gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx := xh.CtxGet()
		defer xh.CtxPut(ctx)

		data, xerr := f(ctx)
		if xerr != nil {
			fmt.Println("err:", xerr)
		} else {
			fmt.Println("succ:", data)
		}

		g.JSON(http.StatusOK, xresponse.JSON(ctx, data, xerr))
		return
	}
}
