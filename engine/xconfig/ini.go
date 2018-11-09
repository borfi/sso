package xconfig

import (
	"flag"
	"log"

	"github.com/robfig/config"
)

// Config 配置单元
type Config struct {
	configFile string
	configs    *config.Config
}

var xini *Config

// New 实例化
func New() *Config {
	if xini != nil {
		return xini
	}

	configName := flag.String("config", "etc/config.ini", "General configuration file")
	flag.Parse() //解析输入的参数

	xini = &Config{}
	xini.configFile = *configName
	xini.Analysis()

	return xini
}

// Analysis 解析配置
func (xf *Config) Analysis() error {
	f, err := config.ReadDefault(xf.configFile)
	if err != nil {
		log.Fatalf("unable to read target config file '%v', err: %v", xf.configFile, err)
		return err
	}
	xf.configs = f
	return err
}

// String 取字符串型
func (xf *Config) String(section, key string) (string, error) {
	return xf.configs.String(section, key)
}

// Int 取整型
func (xf *Config) Int(section, key string) (int, error) {
	return xf.configs.Int(section, key)
}

// Float  取浮点型
func (xf *Config) Float(section, key string) (float64, error) {
	return xf.configs.Float(section, key)
}

// Bool 取布尔型
func (xf *Config) Bool(section, key string) (bool, error) {
	return xf.configs.Bool(section, key)
}
