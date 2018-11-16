package main

import (
	_ "net/http/pprof"
	"sso/engine/xengine"
)

func main() {
	xengine.Run()
}
