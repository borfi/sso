package xengine

import (
	"sso/xengine/xdefine"
	"sync"
)

// 引擎结构体
type xEngine struct {
	lock   sync.RWMutex //引擎锁
	status bool         //引擎状态 true:启动 false:关闭
	close  chan bool    //引擎关闭开关
}

var (
	xengine *xEngine //引擎对象
)

// Get 获取引擎控制器
func Get() xdefine.Engine {
	return xengine
}

// New 启动引擎（引擎必须启动才能工作）
func New() xdefine.Engine {
	if xengine != nil {
		return xengine
	}

	xengine = &xEngine{
		status: true,
		close:  make(chan bool),
	}

	return xengine
}

// Wait 等待引擎关闭
func (x *xEngine) Wait() {
	<-x.close
}

// Close 关闭引擎
func (x *xEngine) Close() {
	x.status = false
	x.close <- true
}

// Status 查看引擎开启状态
func (x *xEngine) Status() bool {
	return x.status
}
