package run

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sso/tools/utils"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Http .
func Http(serverPort int, router *gin.Engine) {
	ip, _ := utils.GetIP()
	addr := fmt.Sprintf("%s:%v", ip, serverPort)
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("start listen server: http://%v", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen err: %s", err)
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
