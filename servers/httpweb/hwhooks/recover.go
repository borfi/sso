package hwhooks

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"sso/xengine/xrecovery"

	"github.com/gin-gonic/gin"
)

// Recovery 拦截panic
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := xrecovery.Stack(3)
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				fmt.Printf("[Recovery] %s\n %s\n", string(httprequest), string(stack))

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
