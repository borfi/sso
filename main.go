package main

import (
	_ "net/http/pprof"
	"sso/engine/xengine"
	"sso/services"
)

func main() {
	xengine.Run(services.Register())
}
