package xserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	readTimeout          time.Duration = 1           // 读超时时间
	writeTimeout         time.Duration = 1           // 写超时时间
	gracefullExitTimeout time.Duration = 1           // 优雅关闭的超时时间
	maxHeaderBytes       int           = 1 << 20     // 请求的头域最大允许长度 1M
	serverName                         = "XEngine"   //服务名
	serverIP             string        = "127.0.0.1" //服务ip
	serverPort                         = 8080        //服务端口
)

// HTTP 服务
type HTTP struct {
	config Config       //配置
	server *http.Server //服务
	engine *gin.Engine  //引擎实例
}

// Config 配置
type Config struct {
	ServerName           string        //服务名
	ServerIP             string        //服务ip
	ServerPort           int           //端口
	ReadTimeout          time.Duration // 读超时时间
	WriteTimeout         time.Duration // 写超时时间
	MaxHeaderBytes       int           // 请求的头域最大允许长度 1M
	GracefullExitTimeout time.Duration // 优雅关闭的超时时间
}

// NewHTTP 新建服务
func NewHTTP() *HTTP {
	xh := &HTTP{}

	// mode
	gin.SetMode(gin.ReleaseMode)

	// disable
	gin.DisableConsoleColor()

	r := gin.New()

	xh.engine = r

	return xh
}

// Run 运行服务
func (xh *HTTP) Run(conf Config) {
	if xh.engine == nil {
		panic("engine not initialized")
	}

	xh.initConfig(conf)
	xh.configServer()

	go func() {
		if err := xh.listen(); err != nil {
			panic(err)
		}
	}()
}

// Hooks 注册中间件
func (xh *HTTP) Hooks(hooks ...gin.HandlerFunc) {
	xh.engine.Use(hooks...)
}

// LoadHTMLGlob 模板渲染
func (xh *HTTP) LoadHTMLGlob(path string) {
	xh.engine.LoadHTMLGlob(path)
}

// Router 设置路由
func (xh *HTTP) Router(f func(*gin.Engine)) {
	if xh.engine == nil {
		return
	}

	f(xh.engine)
}

// Name 取服务名称
func (xh *HTTP) Name() string {
	return xh.config.ServerName
}

// GracefullyExit 完美退出
func (xh *HTTP) GracefullyExit() {
	if xh.server == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), xh.config.GracefullExitTimeout*time.Second)
	defer cancel()

	if err := xh.server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown err: %v", err)
	}
}

// 初始化配置
func (xh *HTTP) initConfig(conf Config) {
	if conf.ServerName == "" {
		conf.ServerName = serverName
	}

	if conf.ServerIP == "" {
		conf.ServerIP = serverIP
	}

	if conf.ServerPort <= 0 {
		conf.ServerPort = serverPort
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

	if conf.GracefullExitTimeout <= 0 {
		conf.GracefullExitTimeout = gracefullExitTimeout
	}

	xh.config = conf
}

// 配置服务
func (xh *HTTP) configServer() {
	xh.server = &http.Server{
		Addr:           xh.listenAddr(),
		Handler:        xh.engine,
		ReadTimeout:    xh.config.ReadTimeout * time.Second,
		WriteTimeout:   xh.config.WriteTimeout * time.Second,
		MaxHeaderBytes: xh.config.MaxHeaderBytes,
	}
}

// 取监听地址
func (xh *HTTP) listenAddr() string {
	return fmt.Sprintf("%s:%d", xh.config.ServerIP, xh.config.ServerPort)
}

// 开始监听
func (xh *HTTP) listen() error {
	if err := xh.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}
