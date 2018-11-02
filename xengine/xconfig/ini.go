package xconfig

import (
	"flag"
	"log"

	"github.com/robfig/config"
)

var myconf *Xconfig

// Xconfig .
type Xconfig struct {
	configFile string
	configs    *config.Config
}

// Config 实例化
func Config() *Xconfig {
	if myconf == nil {
		configName := flag.String("config", "etc/config.ini", "General configuration file")
		flag.Parse() //解析输入的参数

		myconf = &Xconfig{}
		myconf.configFile = *configName
		myconf.Analysis()
	}
	return myconf
}

// Analysis 解析配置
func (xf *Xconfig) Analysis() error {
	f, err := config.ReadDefault(myconf.configFile)
	if err != nil {
		log.Fatalf("Unable to read target config file '%v', err: %v", myconf.configFile, err)
		return err
	}
	myconf.configs = f
	return err
}

// String 取字符串型
func (xf *Xconfig) String(section, key string) (string, error) {
	return xf.configs.String(section, key)
}

// Int 取整型
func (xf *Xconfig) Int(section, key string) (int, error) {
	return xf.configs.Int(section, key)
}

// Float  取浮点型
func (xf *Xconfig) Float(section, key string) (float64, error) {
	return xf.configs.Float(section, key)
}

// Bool 取布尔型
func (xf *Xconfig) Bool(section, key string) (bool, error) {
	return xf.configs.Bool(section, key)
}
