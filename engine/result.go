package engine

import "github.com/gin-gonic/gin"

// JSON .
func JSON(c *gin.Context, code int) {
	r, ok := myEngine.codes[code]
	if !ok {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "未找到匹配的返回码",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	})
}
