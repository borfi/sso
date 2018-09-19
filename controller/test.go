package controller

import (
	"log"
	"sso/code"
	"sso/engine"
	"sso/engine/xconfig"
	"time"

	"github.com/gin-gonic/gin"
)

//Test ...
func Test(c *gin.Context) {
	go func() {
		log.Println("11", c.Query("name"))
		time.Sleep(10 * time.Second)
		log.Println("22", c.Query("name"))
	}()
	port, _ := xconfig.Config().String("service", "port")
	engine.JSON(c, code.Success, port)
}

//Test2 ...
func Test2(c *gin.Context) {
	engine.JSON(c, code.AnalysisConfigError, nil)
}
