package xhttp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// XHTTP 服务
type XHTTP struct {
	core   *gin.Engine
	server *http.Server
}

// New 新建一个http服务
func New(s *http.Server) *XHTTP {
	return &XHTTP{
		core:   gin.New(),
		server: s,
	}
}

// Server 取服务配置
func (xh *XHTTP) Server() *http.Server {
	return xh.server
}

// Listen 监听端口
func (xh *XHTTP) Listen(port int) {
	if err := xh.server.ListenAndServe(); err != http.ErrServerClosed {
		panic(fmt.Sprintf("HTTP server listen err: %v", err))
	}
}

// WebHandler web处理器
func (xh *XHTTP) WebHandler(f func(Context) (interface{}, Error)) gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx := x.CtxGet()
		defer x.CtxPut(ctx)

		data, xerr := f(ctx)
		if xerr != nil {
			fmt.Println("err:", xerr)
		} else {
			fmt.Println("succ:", data)
		}

		g.JSON(http.StatusOK, responseJSON(ctx, data, xerr))
		return
	}
}

// APIHandler api处理器
func (xh *XHTTP) APIHandler(f func(Context) (interface{}, Error)) gin.HandlerFunc {
	return func(g *gin.Context) {
		ctx := x.CtxGet()
		defer x.CtxPut(ctx)

		data, xerr := f(ctx)
		if xerr != nil {
			fmt.Println("err:", xerr)
		} else {
			fmt.Println("succ:", data)
		}

		g.JSON(http.StatusOK, responseJSON(ctx, data, xerr))
		return
	}
}
