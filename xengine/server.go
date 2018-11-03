package xengine

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sso/xengine/xdefine"
	"sso/xengine/xhttpweb"
	"sso/xengine/xtcpapi"
	"time"

	"github.com/gin-gonic/gin"
)

// ServerHTTPWeb 新建一个http服务
func (x *xEngine) ServerHTTPWeb(f func(xdefine.Server, *gin.Engine)) xdefine.Server {
	sev := xhttpweb.New()
	router := gin.Default()

	f(sev, router)

	sev.ServerSet(&http.Server{
		Addr:           ":8081",
		Handler:        router,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	sev.Listen()
	sev.GracefullyExit(1 * time.Second)
	return sev
}

// ServerTCPAPI 新建一个tcp服务
func (x *xEngine) ServerTCPAPI() xdefine.Server {
	return xtcpapi.New()
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
			s[i].GracefullyExit(timeout)
		}
	}
	log.Println("Server shutdown end")
}
