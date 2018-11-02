package xengine

import (
	"sync"
)

// 引擎结构体
type xEngine struct {
	lock    sync.RWMutex //引擎锁
	codes   *xCodes      //返回码单元集合
	ctxPool sync.Pool    //context池
	status  bool         //引擎状态 true:启动 false:关闭
	close   chan bool    //引擎关闭开关
}

// XResponse 最终输出给客户端的json结构
type XResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Info string      `json:"info"`
	Data interface{} `json:"data"`
}
