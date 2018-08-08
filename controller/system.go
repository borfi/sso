package controller

import (
	"sso/engine/xconfig"

	"github.com/gin-gonic/gin"
)

//AnalysisConfig 重新解析配置
func AnalysisConfig(c *gin.Context) {
	err := xconfig.Config().Analysis()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 410001,
			"msg":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "update success",
	})
	return
}
