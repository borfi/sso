package engine

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTPService http service
func HTTPService(router *gin.Engine, port int) {
	//ip, _ := xutils.GetIP()
	addr := fmt.Sprintf("%v:%v", "", port)
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20, //1M
	}

	go func() {
		log.Printf("start listen server: http://%v", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen server err: [%v]", err)
		}
	}()

	gracefullyExit(server)
}

// MonitorHTTPService monitor http service
func MonitorHTTPService(port int) {
	addr := fmt.Sprintf("%v:%v", "", port)
	log.Printf("start listen monitor server: http://%v", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("listen monitor server err: [%v]", err)
	}
}

func gracefullyExit(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	//signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Println("shut down server ......")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shut down err: ", err)
	}
	time.Sleep(10 * time.Second)
	log.Println("server exit!")
}
