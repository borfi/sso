package api

import (
	"sso/code"
	"sso/engine/xservice"
)

//Test ...
func Test(ctx xservice.Context) (interface{}, xservice.Error) {
	//port, _ := xconfig.Config().String("service", "port")
	data := `{
		"code":    1,
		"message": "this is a test",
		"info":    "hahaha",
	}`
	return data, ctx.Error(code.AnalysisConfigError, nil)
}
