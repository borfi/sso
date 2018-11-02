package xhttpweb

import (
	"sso/xengine/xdefine"
	"sso/xengine/xerror"
	"time"

	"github.com/gin-gonic/gin"
)

// XContext context
type XContext struct {
	gin.Context
	startTime time.Time
}

// 实例化context
func newContext() xdefine.Context {
	return &XContext{}
}

// Error 组装错误返回
func (ctx *XContext) Error(code int) xdefine.Error {
	//这里可以做多语言版本的翻译处理
	return xerror.New(code)
}
