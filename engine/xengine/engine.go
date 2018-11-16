package xengine

import (
	"os"
	"os/signal"
	"sso/engine/xservice"
	"sso/services"
	"sync"
)

// 引擎结构体
type xEngine struct {
	lock   sync.RWMutex // 引擎锁
	status bool         // 引擎状态 true:运行中 false:已退出
	quit   chan bool    // 引擎退出开关，后台控制时使用
}

var (
	engine *xEngine
)

// Run 启动引擎（引擎必须启动才能工作），启动之前必须要先注册服务，仅能被主程main包调用
func Run() {
	if engine == nil {
		engine = &xEngine{
			status: true,
			quit:   make(chan bool),
		}
	}

	services.Register()

	xservice.Listen()

	waitQuit()
}

// waitQuit 等待引擎退出
func waitQuit() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill) // 进行杀进程

	for {
		select {
		case <-engine.quit:
			engine.status = false
			xservice.GracefullyQuit()
			os.Exit(0)
			break
		case <-quit:
			engine.status = false
			xservice.GracefullyQuit()
			os.Exit(1)
			break
		}
	}
}

// Quit 退出引擎
func Quit() {
	engine.status = false
	engine.quit <- true
}

// Status 查看引擎开启状态
func Status() bool {
	return engine.status
}
