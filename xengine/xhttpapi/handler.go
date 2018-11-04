package xhttpweb

import (
	"net/http"
	"sso/xengine/xdefine"
	"sso/xengine/xresponse"

	"github.com/gin-gonic/gin"
)

// Handler api处理器
func (xh *XHTTP) Handler(f func(xdefine.Context) (interface{}, xdefine.Error)) interface{} {
	var hf gin.HandlerFunc = func(g *gin.Context) {
		ctx := xh.CtxGet()
		defer xh.CtxPut(ctx)

		data, xerr := f(ctx)

		// if xerr != nil {
		// 	log.Printf("err: %d, %s, %s, %v, %v", xerr.Code(), xerr.Msg(), xerr.Info(), xerr.Error())
		// } else {
		// 	log.Printf("succ: %+v", data)
		// }

		g.JSON(http.StatusOK, xresponse.JSON(ctx, data, xerr))
		return
	}
	return hf
}
