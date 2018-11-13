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

// xWeb 服务
type xWeb struct {
	lock          sync.RWMutex // 锁
	currentStatus int          // 当前状态
	config        *WebConfig   // 配置
	service       *http.Server // 服务
	engine        *gin.Engine  // 引擎实例
}

// WebConfig 配置
type WebConfig struct {
	Name                 string            // 服务名
	IP                   string            // 服务ip
	Port                 int               // 监听端口
	ReadTimeout          time.Duration     // 读超时时间
	WriteTimeout         time.Duration     // 写超时时间
	MaxHeaderBytes       int               // 请求的头域最大允许长度 1M
	GracefullQuitTimeout time.Duration     // 优雅退出的超时时间
	Hooks                []gin.HandlerFunc // 注册的钩子
	Router               func(*gin.Engine) // 路由
	TemplatePath         string            // 模板文件目录
}

func newWeb() service {
	return &xWeb{}
}

// 注册服务
func (xh *xWeb) register(conf interface{}) error {
	err := xh.registerConfig(conf)
	if err != nil {
		return err
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	xh.engine = gin.New()
	xh.registerHooks()
	xh.registerRouter()
	xh.registerTemplate()
	xh.registerService()

	return nil
}

// 取服务名称
func (xh *xWeb) getName() string {
	return xh.config.Name
}

// 取服务监听的端口
func (xh *xWeb) getPort() int {
	return xh.config.Port
}

// 开始监听服务
func (xh *xWeb) listen() error {
	if xh.engine == nil {
		return errors.New("engine not initialized")
	}
	if xh.service == nil {
		return errors.New("service not initialized")
	}

	// err != http.ErrServiceClosed
	if err := xh.service.ListenAndServe(); err != nil {
		return err
	}

	xh.service.Close()
	return nil
}

// 取服务当前状态
func (xh *xWeb) getStatus() int {
	return xh.currentStatus
}

// 设置服务当前状态
func (xh *xWeb) setStatus(status int) {
	xh.lock.RLock()
	xh.currentStatus = status
	xh.lock.RUnlock()
}

// 完美退出
func (xh *xWeb) gracefullyQuit() error {
	if xh.service == nil {
		return errors.New("service not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), xh.config.GracefullQuitTimeout*time.Second)
	defer cancel()

	return xh.service.Shutdown(ctx)
}

// 注册配置
func (xh *xWeb) registerConfig(c interface{}) error {
	conf := &WebConfig{}
	if c != nil {
		var ok bool
		if conf, ok = c.(*WebConfig); !ok {
			return errors.New("service config type error")
		}
	}

	if conf.Name == "" {
		conf.Name = serviceName
	}

	if conf.Port <= 0 {
		conf.Port = servicePort
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
	return nil
}

// 注册钩子
func (xh *xWeb) registerHooks() {
	if xh.config.Hooks == nil || len(xh.config.Hooks) == 0 {
		return
	}
	xh.engine.Use(xh.config.Hooks...)
}

// 注册路由
func (xh *xWeb) registerRouter() {
	if xh.config.Router == nil {
		return
	}
	xh.config.Router(xh.engine)
}

// 注册渲染模板
func (xh *xWeb) registerTemplate() {
	if xh.config.TemplatePath == "" {
		return
	}
	xh.engine.LoadHTMLGlob(xh.config.TemplatePath)
}

// 注册服务
func (xh *xWeb) registerService() {
	xh.service = &http.Server{
		Addr:           xh.listenAddr(),
		Handler:        xh.engine,
		ReadTimeout:    xh.config.ReadTimeout * time.Second,
		WriteTimeout:   xh.config.WriteTimeout * time.Second,
		MaxHeaderBytes: xh.config.MaxHeaderBytes,
	}
}

// 取监听地址
func (xh *xWeb) listenAddr() string {
	return fmt.Sprintf("%s:%d", xh.config.IP, xh.config.Port)
}
