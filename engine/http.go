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

// HTTPService http service
func HTTPService(router *gin.Engine, port int) {
	//ip, _ := xutils.GetIP()

	idleConnsClosed := make(chan struct{})

	addr := getListenAddr(port)

	server := getServerConfig(router, addr)

	go gracefullyExit(server, idleConnsClosed)

	log.Printf("Start HTTP server listen: %v", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server listen: [%v]", err)
	}

	<-idleConnsClosed
}

// HTTPMonitorService monitor http service
func HTTPMonitorService(port int) {
	addr := getListenAddr(port)

	log.Printf("Start HTTP monitor server listen: %v", port)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("HTTP monitor server listen: [%v]", err)
	}
}

// get server config
func getServerConfig(router *gin.Engine, addr string) *http.Server {
	return &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20, //1M
	}
}

// get listen addr
func getListenAddr(port int) string {
	return fmt.Sprintf("%v:%v", "", port)
}

// grace fully exit
func gracefullyExit(server *http.Server, idleConnsClosed chan struct{}) {
	quit := make(chan os.Signal, 1)

	//signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("HTTP server shutdown ......")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown: [%v]", err)
	}

	close(idleConnsClosed)
}
