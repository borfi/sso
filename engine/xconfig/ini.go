package xconfig

import (
	"flag"
	"log"

	"github.com/robfig/config"
)

// Xconfig .
type Xconfig struct {
	configFile string
	configs    *config.Config
}

var myconf *Xconfig

// Config 实例化
func Config() *Xconfig {
	if myconf == nil {
		configName := flag.String("config", "etc/config.ini", "General configuration file")
		flag.Parse() //解析输入的参数

		myconf = &Xconfig{}
		myconf.configFile = *configName
		myconf.analysisIni()
	}
	return myconf
}

func (xf *Xconfig) analysisIni() {
	f, err := config.ReadDefault(myconf.configFile)
	if err != nil {
		log.Fatalf("Unable to read target config file '%v', err: %v", myconf.configFile, err)
		return
	}
	myconf.configs = f
}

// String .
func (xf *Xconfig) String(section, key string) (string, error) {
	return xf.configs.String(section, key)
}

// Int .
func (xf *Xconfig) Int(section, key string) (int, error) {
	return xf.configs.Int(section, key)
}

// Float  .
func (xf *Xconfig) Float(section, key string) (float64, error) {
	return xf.configs.Float(section, key)
}

// Bool .
func (xf *Xconfig) Bool(section, key string) (bool, error) {
	return xf.configs.Bool(section, key)
}
