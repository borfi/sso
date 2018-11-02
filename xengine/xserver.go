package xengine

import (
	"sso/xengine/xhttp"
	"sso/xengine/xtcp"
)

// Server .
type Server interface {
}

// NewHTTP 新建一个http服务
func (x *xEngine) NewHTTP() Server {
	return &xhttp.New()
}

// NewTCP 新建一个tcp服务
func (x *xEngine) NewTCP() Server {
	return &xtcp.New()
}
