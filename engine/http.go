package engine

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sso/engine/xconfig"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	readTimeout             time.Duration = 1       // 读超时时间
	writeTimeout            time.Duration = 3       // 写超时时间
	gracefullExitCtxTimeout time.Duration = 5       // 优雅关闭的超时时间
	maxHeaderBytes          int           = 1 << 20 // 请求的头域最大允许长度 1M
)

// RunHTTPService http service
func RunHTTPService(router *gin.Engine) {
	//ip, _ := xutils.GetIP()

	httpServiceClosed := make(chan struct{})

	port, err := getHTTPServicePort()
	if err != nil {
		panic(err.Error())
	}

	addr := getListenAddr(port)
	server := getServerConfig(router, addr)

	go gracefullyExit(server, httpServiceClosed)

	log.Printf("Start http server listen: %v", port)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(fmt.Sprintf("HTTP server listen err: %v", err))
	}

	<-httpServiceClosed
}

// RunHTTPMonitorService monitor http service
func RunHTTPMonitorService() {
	port, err := getHTTPMonitorServicePort()
	if err != nil {
		panic(err.Error())
	}

	addr := getListenAddr(port)

	log.Printf("Start http monitor server listen: %v", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(fmt.Sprintf("HTTP monitor server listen err: %v", err))
	}
}

// get service port
func getHTTPServicePort() (int, error) {
	port, err := xconfig.Config().Int("http_service", "port")
	if err != nil {
		log.Fatalf("Get http service port err: %v", err)
		return 0, err
	}
	return port, err
}

// get monitor port
func getHTTPMonitorServicePort() (int, error) {
	port, err := xconfig.Config().Int("http_service", "monitor_port")
	if err != nil {
		log.Fatalf("Get http monitor service port err: %v", err)
		return 0, err
	}
	return port, err
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
func gracefullyExit(server *http.Server, httpServiceClosed chan struct{}) {
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

	close(httpServiceClosed)
}
