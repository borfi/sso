package engine

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// JSON .
func JSON(c *gin.Context, code int, data interface{}) {
	//execute time
	t := c.GetTime("sso_exec_start_time")
	execTime := time.Now().Sub(t)

	//Not find code
	r, ok := myEngine.codes[code]
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":     0,
			"msg":      "No matching return code was found",
			"data":     nil,
			"exectime": fmt.Sprintf("%13v", execTime),
		})
		return
	}

	//normal return
	c.JSON(http.StatusOK, gin.H{
		"code":     r.Code,
		"msg":      r.Msg,
		"data":     data,
		"exectime": fmt.Sprintf("%v", execTime),
	})
}
