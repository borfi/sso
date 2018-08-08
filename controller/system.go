package controller

import (
	"sso/engine"
	"sso/engine/xconfig"
	"sso/errorcode"

	"github.com/gin-gonic/gin"
)

//AnalysisConfig 重新解析配置
func AnalysisConfig(c *gin.Context) {
	err := xconfig.Config().Analysis()
	if err != nil {
		engine.JSON(c, errorcode.AnalysisConfigError)
		return
	}
	engine.JSON(c, errorcode.Success)
	return
}
