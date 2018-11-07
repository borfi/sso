package hwcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Test .
func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "site/index.html", gin.H{
		"title":   "Main website",
		"content": "Hello engine!~",
	})
}

// //TestSessionSet ...
// func TestSessionSet(c *gin.Context) {
// 	session := sessions.Default(c)

// 	session.Set("toto", "hahahahahhahahahha")
// 	session.Set("dodo", 1234556)
// 	session.Save()

// 	engine.JSON(c, code.Success, nil)
// }

// //TestSessionGet ...
// func TestSessionGet(c *gin.Context) {
// 	session := sessions.Default(c)

// 	toto := session.Get("toto")
// 	dodo := session.Get("dodo")
// 	eoeo := session.Get("eoeo")

// 	data := []interface{}{
// 		toto, dodo, eoeo,
// 	}

// 	engine.JSON(c, code.Success, data)
// }
