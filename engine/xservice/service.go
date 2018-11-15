package xservice

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Context 服务上下文接口
type Context interface {
	context.Context
	Error(int, error) Error
}

// Error 错误接口
type Error interface {
	GetCode() int                              // 取错误码
	GetMsg() string                            // 取给外界看的错误信息
	GetInfo() string                           // 取给内部看的错误信息
	GetError() error                           // 取具体的错误
	GetFields() map[string]interface{}         // 取相关的数据
	SetFields(fs map[string]interface{}) Error // 设置相关数据
	Format(fields []interface{}) Error         // 格式化错误信息（内部+外界）
}

// service 服务接口
type service interface {
	register(interface{}) error // 注册服务
	listen() error              // 开始监听
	gracefullyQuit() error      // 完美退出
	getName() string            // 取服务名称
	getPort() int               // 取监听的端口
	setStatus(int)              // 设置服务当前状态
	getStatus() int             // 取服务当前状态
}

// Router 单个路由配置
type Router struct {
	Method  string
	Path    string
	Handler Handler
}

// Handler 处理器
type Handler func(Context) (interface{}, Error)

// 服务控制器
type serviceController struct {
	lock     sync.RWMutex          // 锁
	all      []service             // 所有的服务列表
	stopping chan *stoppingService // 正在停止的服务队列
}

// 停止监听的服务队列单元
type stoppingService struct {
	serv service
	err  error
}

// 最终输出的json结构
type response struct {
	Code int         `json:"code"` // 返回码
	Msg  string      `json:"msg"`  // 给外界看的信息
	Info string      `json:"info"` // 给内部看的信息
	Data interface{} `json:"data"` // 数据
}

const (
	readTimeout          time.Duration = 1         // 读超时时间
	writeTimeout         time.Duration = 1         // 写超时时间
	gracefullQuitTimeout time.Duration = 1         // 优雅退出的超时时间
	serviceName          string        = "XEngine" // 服务名
	servicePort          int           = 8080      // 服务端口
	maxHeaderBytes       int           = 1 << 20   // 请求头的最大允许长度 1M
	maxRequestBodySize   int           = 1 << 20   // 请求体的最大允许长度 1M

	// 各种服务当前状态
	statusInit       int = 0 // 初始化状态
	statusRunning    int = 1 // 运行中状态
	statusListenFail int = 2 // 监听失败(没有启动成功，正常看不到这个状态，因为监听失败要立即panic)
	statusQuit       int = 3 // 退出状态
)

var (
	servicer *serviceController
)

func init() {
	servicer = &serviceController{
		all:      make([]service, 0),
		stopping: make(chan *stoppingService),
	}
}

// RegisterHTTP 注册服务
func RegisterHTTP(conf *HTTPConfig) {
	xs := newHTTP()
	register(xs, conf)
}

// RegisterWeb 注册服务
func RegisterWeb(conf *WebConfig) {
	xs := newWeb()
	register(xs, conf)
}

// 注册通用处理
func register(xs service, conf interface{}) {
	err := xs.register(conf)
	if err != nil {
		panic(err)
	}

	if isListened(xs.getPort()) {
		panic(fmt.Errorf("port %d has been occupied", xs.getPort()))
	}

	servicer.all = append(servicer.all, xs)
}

// Count 服务总数
func Count() int {
	return len(servicer.all)
}

// Listen 开始监听服务
func Listen() {
	for i := 0; i < len(servicer.all); i++ {
		if !isInitStatus(servicer.all[i].getStatus()) {
			continue
		}

		go func(i int) {
			if err := servicer.all[i].listen(); err != nil {
				if err != http.ErrServerClosed {
					servicer.all[i].setStatus(statusListenFail)
					panic(fmt.Sprintf("service %s %d listen error, %v", servicer.all[i].getName(), servicer.all[i].getPort(), err))
				}

				servicer.all[i].setStatus(statusQuit)
				servicer.stopping <- &stoppingService{
					serv: servicer.all[i],
					err:  err,
				}
				return
			}

			servicer.all[i].setStatus(statusRunning)
		}(i)
	}
}

// GracefullyQuit 完美退出
func GracefullyQuit() {
	if len(servicer.all) == 0 {
		return
	}

	log.Println("sevice shutdown start")
	for i := 0; i < len(servicer.all); i++ {
		if err := servicer.all[i].gracefullyQuit(); err != nil {
			log.Printf("sevice %s gracefully quit error %v", servicer.all[i].getName(), err)
			continue
		}
		log.Printf("sevice %s gracefully quit", servicer.all[i].getName())
	}
	log.Println("sevice shutdown end")
}

// 端口是否已经被监听
func isListened(port int) bool {
	if len(servicer.all) == 0 {
		return false
	}

	for _, v := range servicer.all {
		if v.getPort() == port {
			return true
		}
	}
	return false
}

// isInitStatus 判断是否为初始化状态
func isInitStatus(status int) bool {
	return status == statusInit
}

// isRunningStatus 判断是否为运行状态
func isRunningStatus(status int) bool {
	return status == statusRunning
}

// isListenFailStatus 判断是否为监听失败状态
func isListenFailStatus(status int) bool {
	return status == statusListenFail
}

// isQuitStatus 判断是否为退出状态
func isQuitStatus(status int) bool {
	return status == statusQuit
}
