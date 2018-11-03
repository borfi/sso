package xdefine

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Engine 引擎控制器
type Engine interface {
	Status() bool                                   //返回引擎状态
	Wait()                                          //阻塞并等待引擎被关闭
	Close()                                         //关闭引擎
	ServerHTTPWeb(func(Server, *gin.Engine)) Server //创建一个http服务
	ServerTCPAPI() Server                           //创建一个tcp服务
}

// Server 服务接口
type Server interface {
	Name() string                                           //服务名
	ServerSet(*http.Server)                                 //设置服务配置
	Handler(func(Context) (interface{}, Error)) interface{} //服务处理器
	Listen()                                                //开始监听
	CtxGet() Context                                        //从池子里取出context
	CtxPut(ctx Context)                                     //返还context到池子
	GracefullyExit(time.Duration)                           //完美退出
}

// Context 上下文接口
type Context interface {
	Error(int) Error //返回错误
}

// Error 错误接口
type Error interface {
	Code() int                   //返回错误码
	Msg() string                 //返回可以给外界看的信息
	Info() string                //返回可以给内部看的信息
	SetError(error) Error        //设置错误
	Format(...interface{}) Error //格式化错误信息（内部+外界）
}
