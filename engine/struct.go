package engine

import (
	"net/http"
	"sync"
	"time"
)

// 引擎结构体
type xEngine struct {
	lock    sync.RWMutex //引擎锁
	codes   *xCodes      //返回码单元集合
	ctxPool sync.Pool    //context池
	status  bool         //引擎状态 true:启动 false:关闭
	close   chan bool    //引擎关闭开关
}

// XContext context
type XContext struct {
	startTime time.Time
	request   *http.Request
	response  *http.Response
}

// 错误信息
type xError struct {
	code int
	msg  string
	info string
	err  error
}

// 返回码总单元
type xCodes struct {
	m sync.Map
}

// XCode 返回码单元
type XCode struct {
	Code int    //返回码
	Msg  string //给外界看的信息
	Info string //给开发人员看的信息
}

// XResponse 最终输出给客户端的json结构
type XResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Info string      `json:"info"`
	Data interface{} `json:"data"`
}
