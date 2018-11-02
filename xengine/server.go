package xengine

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sso/xengine/xdefine"
	"sso/xengine/xhttpweb"
	"sso/xengine/xresponse"
	"time"

	"github.com/gin-gonic/gin"
)

// ServerHTTP 新建一个http服务
func (x *xEngine) ServerHTTP(f func(*gin.Engine)) xdefine.Server {
	sev := &xhttpweb.New()
	app := gin.Default()
	f(app)
	sev.Listen()
	sev.GracefullyExit(1 * time.Second)
	return sev
}

// ServerTCP 新建一个tcp服务
func (x *xEngine) ServerTCP() xdefine.Server {
	return &xtcp.New()
}

// Handler 处理器
func (x *xEngine) Handler(f func(xdefine.Context) (interface{}, xdefine.Error)) gin.HandlerFunc {
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

// GracefullyExit 完美退出
func (x *xEngine) GracefullyExit(timeout time.Duration, s ...xdefine.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Println("Server shutdown start")
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			log.Printf("Server name: %s", s[i].Name())
			s[i].GraceFullyExit(timeout)
		}
	}
	log.Println("Server shutdown end")
}
