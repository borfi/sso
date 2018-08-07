package xconfig

import (
	"flag"
	"log"

	"github.com/robfig/config"
)

// Xconfig .
type Xconfig struct {
	ConfigFile string
	ConfigMap  map[string]string
}

var myconf *Xconfig

// Xconf 实例化
func Xconf() *Xconfig {
	if myconf == nil {
		myconf = &Xconfig{}
		myconf.ConfigFile = flag.String("f", "/etc/config.ini", "General configuration file")
		myconf.NewIni()
	}
	return myconf
}

// NewIni .
func (xf *Xconfig) NewIni() {
	f, error := config.ReadDefault(f.ConfigFile)
	if error != nil {
		log.Fatalf("Unable to read target config file '%s'", xf.ConfigFile)
	}
	xf.ConfigMap["ip"] = f.String("service", "ip")
	xf.ConfigMap["port"] = f.String("service", "port")
	xf.ConfigMap["name"] = f.String("service", "name")
}

// ServerPort 获取服务端口
func (xf *Xconfig) ServerPort() (port int) {
	port = xf.ConfigMap["port"]
	return
}
