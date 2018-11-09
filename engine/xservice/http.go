package xservice

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// xHTTP 服务
type xHTTP struct {
	lock          sync.RWMutex // 锁
	currentStatus int          // 当前状态
	config        *ConfigHTTP  // 配置
	service       *http.Server // 服务
	engine        *gin.Engine  // 引擎实例
}

// ConfigHTTP 配置
type ConfigHTTP struct {
	ServiceName          string            // 服务名
	ServiceIP            string            // 服务ip
	ServicePort          int               // 监听端口
	ReadTimeout          time.Duration     // 读超时时间
	WriteTimeout         time.Duration     // 写超时时间
	MaxHeaderBytes       int               // 请求的头域最大允许长度 1M
	GracefullQuitTimeout time.Duration     // 优雅退出的超时时间
	Hooks                []gin.HandlerFunc //注册的钩子
	Router               func(*gin.Engine) //路由
	TemplatePath         string            //模板文件目录
}

var (
	errConfigType = errors.New("service config type error")
	errEngineNil  = errors.New("engine not initialized")
	errServiceNil = errors.New("service not initialized")
)

func newHTTP() service {
	return &xHTTP{}
}

// 注册服务
func (xh *xHTTP) register(conf interface{}) error {
	c, ok := conf.(*ConfigHTTP)
	if !ok {
		return errConfigType
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	xh.engine = gin.New()
	xh.registerConfig(c) // must first register
	xh.registerHooks()
	xh.registerRouter()
	xh.registerTemplate()
	xh.registerService()

	return nil
}

// 取服务名称
func (xh *xHTTP) name() string {
	return xh.config.ServiceName
}

// 取服务监听的端口
func (xh *xHTTP) port() int {
	return xh.config.ServicePort
}

// 开始监听服务
func (xh *xHTTP) listen() error {
	if xh.engine == nil {
		return errEngineNil
	}
	if xh.service == nil {
		return errServiceNil
	}

	// err != http.ErrServiceClosed
	if err := xh.service.ListenAndServe(); err != nil {
		return err
	}

	xh.service.Close()
	return nil
}

// 取服务当前状态
func (xh *xHTTP) status() int {
	return xh.currentStatus
}

// 设置服务当前状态
func (xh *xHTTP) setStatus(status int) {
	xh.lock.RLock()
	xh.currentStatus = status
	xh.lock.RUnlock()
}

// 完美退出
func (xh *xHTTP) gracefullyQuit() error {
	if xh.service == nil {
		return errServiceNil
	}

	ctx, cancel := context.WithTimeout(context.Background(), xh.config.GracefullQuitTimeout*time.Second)
	defer cancel()

	return xh.service.Shutdown(ctx)
}

// 注册配置
func (xh *xHTTP) registerConfig(conf *ConfigHTTP) {
	if conf == nil {
		conf = &ConfigHTTP{}
	}

	if conf.ServiceName == "" {
		conf.ServiceName = serviceName
	}

	if conf.ServiceIP == "" {
		conf.ServiceIP = serviceIP
	}

	if conf.ServicePort <= 0 {
		conf.ServicePort = servicePort
	}

	if conf.ReadTimeout <= 0 {
		conf.ReadTimeout = readTimeout
	}

	if conf.WriteTimeout <= 0 {
		conf.WriteTimeout = writeTimeout
	}

	if conf.MaxHeaderBytes <= 0 {
		conf.MaxHeaderBytes = maxHeaderBytes
	}

	if conf.GracefullQuitTimeout <= 0 {
		conf.GracefullQuitTimeout = gracefullQuitTimeout
	}

	xh.config = conf
}

// 注册钩子
func (xh *xHTTP) registerHooks() {
	if xh.config.Hooks == nil || len(xh.config.Hooks) == 0 {
		return
	}
	xh.engine.Use(xh.config.Hooks...)
}

// 注册路由
func (xh *xHTTP) registerRouter() {
	if xh.config.Router == nil {
		return
	}
	xh.config.Router(xh.engine)
}

// 注册渲染模板
func (xh *xHTTP) registerTemplate() {
	if xh.config.TemplatePath == "" {
		return
	}
	xh.engine.LoadHTMLGlob(xh.config.TemplatePath)
}

// 注册服务
func (xh *xHTTP) registerService() {
	xh.service = &http.Server{
		Addr:           xh.listenAddr(),
		Handler:        xh.engine,
		ReadTimeout:    xh.config.ReadTimeout * time.Second,
		WriteTimeout:   xh.config.WriteTimeout * time.Second,
		MaxHeaderBytes: xh.config.MaxHeaderBytes,
	}
}

// 取监听地址
func (xh *xHTTP) listenAddr() string {
	return fmt.Sprintf("%s:%d", xh.config.ServiceIP, xh.config.ServicePort)
}
