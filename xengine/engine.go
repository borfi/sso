package xengine

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

// Engine 引擎控制器
type Engine interface {
	Status() bool //返回引擎状态
	Wait()        //阻塞并等待引擎被关闭
	Close()       //关闭引擎
}

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
func Get() Engine {
	return xengine
}

// New 启动引擎（引擎必须启动才能工作）
func New() Engine {
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

// GracefullyExit 完美退出
func (x *xEngine) GracefullyExit(s ...define.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	log.Println("Server shutdown start")
	if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			log.Printf("Server name: %s", s[i].Name())
			s[i].GracefullyExit()
		}
	}
	log.Println("Server shutdown end")
}
