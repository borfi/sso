package main

import (
	_ "net/http/pprof"
	"sso/engine/xengine"
	service "sso/services"
)

func main() {
	xengine.Run(service.Register())
}
