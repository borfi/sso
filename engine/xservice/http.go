package xservice

import (
	"errors"
	"fmt"
	"sso/engine/xcode"
	"sso/engine/xutils"
	"sync"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// HTTPConfig 服务配置
type HTTPConfig struct {
	Name                 string        // 服务名
	IP                   string        // 服务ip
	Port                 int           // 监听端口
	ReadTimeout          time.Duration // 读超时时间
	WriteTimeout         time.Duration // 写超时时间
	MaxRequestBodySize   int           // 请求体的最大允许长度 默认1M
	GracefullQuitTimeout time.Duration // 优雅退出的超时时间
	Hooks                []Handler     // 注册的钩子
	Router               []*Router     // 路由
}

// xHTTP 服务
type xHTTP struct {
	lock    sync.RWMutex     // 锁
	status  int              // 当前状态
	config  *HTTPConfig      // 配置
	service *fasthttp.Server // 服务
	ctx     sync.Pool        // ctx连接池
}

func newHTTP() service {
	return &xHTTP{
		ctx: sync.Pool{
			New: func() interface{} {
				return &xContext{}
			},
		},
	}
}

// 注册服务
func (xh *xHTTP) register(conf interface{}) error {
	err := xh.registerConfig(conf)
	if err != nil {
		return err
	}

	err = xh.registerService()
	return err
}

// 取服务名称
func (xh *xHTTP) getName() string {
	return xh.config.Name
}

// 取服务监听的端口
func (xh *xHTTP) getPort() int {
	return xh.config.Port
}

// 开始监听服务
func (xh *xHTTP) listen() error {
	if xh.service == nil {
		return errors.New("service not initialized")
	}

	if err := xh.service.ListenAndServe(xh.listenAddr()); err != nil {
		return err
	}
	return nil
}

// 取服务当前状态
func (xh *xHTTP) getStatus() int {
	return xh.status
}

// 设置服务当前状态
func (xh *xHTTP) setStatus(status int) {
	xh.lock.RLock()
	xh.status = status
	xh.lock.RUnlock()
}

// 完美退出
func (xh *xHTTP) gracefullyQuit() error {
	if xh.service == nil {
		return errors.New("service not initialized")
	}

	return xh.service.Shutdown()
}

// 注册配置
func (xh *xHTTP) registerConfig(c interface{}) error {
	conf := &HTTPConfig{}
	if c != nil {
		var ok bool
		if conf, ok = c.(*HTTPConfig); !ok {
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

	if conf.MaxRequestBodySize <= 0 {
		conf.MaxRequestBodySize = maxRequestBodySize
	}

	if conf.GracefullQuitTimeout <= 0 {
		conf.GracefullQuitTimeout = gracefullQuitTimeout
	}

	xh.config = conf
	return nil
}

// 注册服务
func (xh *xHTTP) registerService() error {
	router, err := xh.registerRouter()
	if err != nil {
		return err
	}

	xh.service = &fasthttp.Server{
		Handler:            router.Handler,
		ReadTimeout:        xh.config.ReadTimeout * time.Second,
		WriteTimeout:       xh.config.WriteTimeout * time.Second,
		MaxRequestBodySize: xh.config.MaxRequestBodySize,
	}
	return nil
}

// 注册路由
func (xh *xHTTP) registerRouter() (*fasthttprouter.Router, error) {
	if xh.config.Router == nil || len(xh.config.Router) == 0 {
		return nil, nil
	}

	router := fasthttprouter.New()

	for _, v := range xh.config.Router {
		if !xutils.IsURLPath(v.Path) {
			return nil, errors.New("router path is not allowed")
		}

		switch v.Method {
		case "get", "GET":
			router.GET(v.Path, xh.handler(v.Handler))
			break
		case "post", "POST":
			router.POST(v.Path, xh.handler(v.Handler))
			break
		case "head", "HEAD":
			router.HEAD(v.Path, xh.handler(v.Handler))
			break
		case "options", "OPTIONS":
			router.OPTIONS(v.Path, xh.handler(v.Handler))
			break
		case "put", "PUT":
			router.PUT(v.Path, xh.handler(v.Handler))
			break
		case "patch", "PATCH":
			router.PATCH(v.Path, xh.handler(v.Handler))
			break
		case "delete", "DELETE":
			router.DELETE(v.Path, xh.handler(v.Handler))
			break
		default:
			return nil, errors.New("router method is not allowed")
		}
	}

	return router, nil
}

// 处理器
func (xh *xHTTP) handler(f Handler) fasthttp.RequestHandler {
	return func(fhCtx *fasthttp.RequestCtx) {
		ctx := xh.ctx.Get().(*xContext)
		defer xh.ctx.Put(ctx)
		ctx.fctx = fhCtx

		xerr := xh.runHooks(ctx)
		if xerr != nil {
			// 输出出错结果
			ctx.ResponseJSON(xerr.GetCode(), xerr.GetMsg(), xerr.GetInfo(), nil)
			return
		}

		data, xerr := f(ctx)
		if xerr != nil {
			// 输出出错结果
			ctx.ResponseJSON(xerr.GetCode(), xerr.GetMsg(), xerr.GetInfo(), nil)
			return
		}

		// 输出正确结果
		r := xcode.Get(xcode.Success)
		ctx.ResponseJSON(r.Code, r.Msg, r.Info, data)
		return
	}
}

// 运行hooks
func (xh *xHTTP) runHooks(ctx *xContext) Error {
	if xh.config.Hooks != nil {
		for _, hk := range xh.config.Hooks {
			if hk == nil {
				continue
			}

			_, xerr := hk(ctx)
			if xerr != nil {
				return xerr
			}
		}
	}
	return nil
}

// 取监听地址
func (xh *xHTTP) listenAddr() string {
	return fmt.Sprintf("%s:%d", xh.config.IP, xh.config.Port)
}
