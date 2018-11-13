package xservice

import (
	"context"
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// 上下文
type xContext struct {
	context.Context
	fctx *fasthttp.RequestCtx
}

// Error 生成错误
func (c *xContext) Error(code int, err error) Error {
	return newError(code, err)
}

// ResponseJSON 组装json输出返回
func (c *xContext) ResponseJSON(code int, msg, info string, data interface{}) {
	res := &response{
		Code: code,
		Msg:  msg,
		Info: info,
		Data: data,
	}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		// 这里要记录日志
		return
	}

	// jsonRes = []byte(`{
	// 	"code":    100,
	// 	"message": "解析配置文件失败",
	// 	"info":    "hahaha",
	// }`)

	c.fctx.Write(jsonRes)
	return
}
