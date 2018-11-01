package controller

import (
	"fmt"
	"sso/code"
	"sso/engine"
	"sso/engine/xconfig"
)

//Test ...
func Test(ctx engine.Context) (interface{}, engine.Error) {
	port, _ := xconfig.Config().String("service", "port")
	r := ctx.Error(code.ParamsError).Format(11, "ss")
	fmt.Println(r)
	return port, ctx.Error(code.Success)
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
