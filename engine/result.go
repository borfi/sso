package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON .
func JSON(c *gin.Context, code int, data interface{}) {
	//Not find code
	r := engine.codes.Get(code)

	//normal return
	c.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
		"data": data,
	})
}
