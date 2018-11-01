package engine

import (
	"sync"
)

var (
	engine *XEngine //引擎控制器
)

// XEngine .
type XEngine struct {
	lock    sync.RWMutex //引擎锁
	codes   *XCodes      //返回码
	ctxPool sync.Pool    //context池
	status  bool         //引擎状态 true:启动 false:关闭
	close   chan bool    //引擎关闭开关
}

// Engine 获取引擎控制器
func Engine() *XEngine {
	return engine
}

// 启动引擎（引擎必须启动才能工作）
func start() {
	if engine != nil {
		return
	}

	engine = &XEngine{
		codes: &XCodes{},
		ctxPool: sync.Pool{
			New: func() interface{} {
				return newContext()
			},
		},
		status: true,
		close:  make(chan bool),
	}
}

// Wait 等待引擎关闭
func (x *XEngine) Wait() {
	<-x.close
}

// Close 关闭引擎
func (x *XEngine) Close(ctx *XContext) {
	x.status = false
	x.close <- true
}

// CtxGet 取出context
func (x *XEngine) CtxGet() *XContext {
	return x.ctxPool.Get().(*XContext)
}

// CtxPut 返还context
func (x *XEngine) CtxPut(ctx *XContext) {
	x.ctxPool.Put(ctx)
}

// Status 查看引擎开启状态
func (x *XEngine) Status() bool {
	return x.status
}
