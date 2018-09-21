package engine

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	readTimeout             time.Duration = 1       // 读超时时间
	writeTimeout            time.Duration = 3       // 写超时时间
	gracefullExitCtxTimeout time.Duration = 5       // 优雅关闭的超时时间
	maxHeaderBytes          int           = 1 << 20 // 请求的头域最大允许长度 1M
)

// HTTPService http service
func HTTPService(router *gin.Engine, port int) {
	//ip, _ := xutils.GetIP()

	idleConnsClosed := make(chan struct{})

	addr := getListenAddr(port)

	server := getServerConfig(router, addr)

	go gracefullyExit(server, idleConnsClosed)

	log.Printf("Start http server listen: %v", port)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server listen err: %v", err)
	}

	<-idleConnsClosed
}

// HTTPMonitorService monitor http service
func HTTPMonitorService(port int) {
	addr := getListenAddr(port)

	log.Printf("Start http monitor server listen: %v", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("HTTP monitor server listen err: %v", err)
	}
}

// get server config
func getServerConfig(router *gin.Engine, addr string) *http.Server {
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
}

// get listen addr
func getListenAddr(port int) string {
	return fmt.Sprintf("%v:%v", "", port)
}

// gracefully exit
func gracefullyExit(server *http.Server, idleConnsClosed chan struct{}) {
	quit := make(chan os.Signal, 1)

	//signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("HTTP server shutdown ......")

	ctx, cancel := context.WithTimeout(context.Background(), gracefullExitCtxTimeout*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown err: %v", err)
	}

	close(idleConnsClosed)
}
