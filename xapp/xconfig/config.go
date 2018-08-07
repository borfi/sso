package xconfig

import (
	"flag"
	"log"

	"github.com/robfig/config"
)

// Xconfig .
type Xconfig struct {
	configFile *string
	configs    *config.Config
}

var myconf *Xconfig

// Xconf 实例化
func Xconf() *Xconfig {
	if myconf == nil {
		myconf = &Xconfig{}
		myconf.configFile = flag.String("f", "etc/config.ini", "General configuration file")
		myconf.analysisIni()
	}
	return myconf
}

func (xf *Xconfig) analysisIni() {
	f, err := config.ReadDefault(*myconf.configFile)
	if err != nil {
		log.Fatalf("Unable to read target config file '%v', err: %v", *myconf.configFile, err)
		return
	}
	myconf.configs = f
}

// String .
func (xf *Xconfig) String(section, option string) (string, error) {
	return xf.configs.String(section, option)
}

// Int .
func (xf *Xconfig) Int(section, option string) (int, error) {
	return xf.configs.Int(section, option)
}

// Float  .
func (xf *Xconfig) Float(section, option string) (float64, error) {
	return xf.configs.Float(section, option)
}

// Bool .
func (xf *Xconfig) Bool(section, option string) (bool, error) {
	return xf.configs.Bool(section, option)
}
