package xrun

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sso/xapp/xconfig"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTP .
func HTTP(router *gin.Engine) {
	serverPort, err := xconfig.Xconf().String("service", "port")
	if err != nil {
		log.Fatalf("Get service port err: %v", err)
		return
	}

	//ip, _ := xutils.GetIP()
	addr := fmt.Sprintf("%v:%v", "", serverPort)
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20, //1M
	}

	go func() {
		log.Printf("Start listen server: http://%v", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Listen err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting!")
}
