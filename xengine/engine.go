package xengine

import (
	"sync"
)

var (
	xengine *xEngine //引擎对象
)

// Get 获取引擎控制器
func Get() Engine {
	return xengine
}

// New 启动引擎（引擎必须启动才能工作）
func New() Engine {
	if xengine != nil {
		return xengine
	}

	xengine = &xEngine{
		codes: &xCodes{},
		ctxPool: sync.Pool{
			New: func() interface{} {
				return newContext()
			},
		},
		status: true,
		close:  make(chan bool),
	}

	return xengine
}

// WaitClose 等待引擎关闭
func (x *xEngine) WaitClose() {
	<-x.close
}

// Close 关闭引擎
func (x *xEngine) Close(ctx *XContext) {
	x.status = false
	x.close <- true
}

// CtxGet 取出context
func (x *xEngine) CtxGet() *XContext {
	return x.ctxPool.Get().(*XContext)
}

// CtxPut 返还context
func (x *xEngine) CtxPut(ctx *XContext) {
	x.ctxPool.Put(ctx)
}

// Status 查看引擎开启状态
func (x *xEngine) Status() bool {
	return x.status
}
