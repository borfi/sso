package controller

import (
	"sso/code"
	"sso/engine"
	"sso/engine/xconfig"

	"github.com/gin-gonic/gin"
)

//AnalysisConfig 重新解析配置
func AnalysisConfig(c *gin.Context) {
	err := xconfig.Config().Analysis()
	if err != nil {
		engine.JSON(c, code.AnalysisConfigError, nil)
		return
	}
	engine.JSON(c, code.Success, nil)
	return
}
