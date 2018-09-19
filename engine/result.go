package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON .
func JSON(c *gin.Context, code int, data interface{}) {
	//Not find code
	r, ok := myEngine.codes[code]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "No matching return code was found",
			"data": nil,
		})
		return
	}

	//normal return
	c.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
		"data": data,
	})
}
