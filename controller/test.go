package controller

import (
	"errors"
	"fmt"
	"sso/xengine/xconfig"
	"sso/xengine/xerror"
)

//Test ...
func Test(ctx xengine.Context) (interface{}, xerror.Error) {
	port, _ := xconfig.Config().String("service", "port")

	r := ctx.Error(xcode.ParamsError).Format(11, "ss")
	fmt.Println(r)

	return port, ctx.Error(xcode.AnalysisConfigError).SetError(errors.New("this is test error"))
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
