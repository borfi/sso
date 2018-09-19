package controller

import (
	"log"
	"sso/code"
	"sso/engine"
	"sso/engine/xconfig"
	"time"

	"github.com/gin-contrib/sessions"
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

//TestSessionSet ...
func TestSessionSet(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("toto", "hahahahahhahahahha")
	session.Set("dodo", 1234556)
	session.Save()

	engine.JSON(c, code.Success, nil)
}

//TestSessionGet ...
func TestSessionGet(c *gin.Context) {
	session := sessions.Default(c)

	toto := session.Get("toto")
	dodo := session.Get("dodo")
	eoeo := session.Get("eoeo")

	data := []interface{}{
		toto, dodo, eoeo,
	}

	engine.JSON(c, code.Success, data)
}
