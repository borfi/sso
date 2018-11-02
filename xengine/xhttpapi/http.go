package xhttpweb

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sso/xengine/xdefine"
	"sync"
	"time"
)

// XHTTP 服务
type XHTTP struct {
	name   string       //服务名
	server *http.Server //服务实例
	ctx    sync.Pool    //context池
}

// New 新建一个http服务
func New() xdefine.Server {
	return &XHTTP{
		name: "http web",
		ctx: sync.Pool{
			New: func() interface{} {
				return newContext()
			},
		},
	}
}

// Name 服务名称
func (xh *XHTTP) Name() string {
	return xh.name
}

// ServerSet 服务设置
func (xh *XHTTP) ServerSet(s *http.Server) {
	xh.server = s
}

// Listen 开始监听
func (xh *XHTTP) Listen() {
	if err := xh.server.ListenAndServe(); err != http.ErrServerClosed {
		panic(fmt.Sprintf("HTTP server listen err: %v", err))
	}
}

// CtxGet 取出context
func (xh *XHTTP) CtxGet() xdefine.Context {
	return xh.ctx.Get().(*XContext)
}

// CtxPut 返还context
func (xh *XHTTP) CtxPut(ctx xdefine.Context) {
	xh.ctx.Put(ctx)
}

// GracefullyExit 完美退出
func (xh *XHTTP) GracefullyExit(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	if err := xh.server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown err: %v", err)
	}
}
