package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	logger "github.com/ipfs/go-log"
	"github.com/spf13/viper"
	"os"
	"rcsv/pkg/conf"
	"rcsv/pkg/constant"
	"rcsv/pkg/utils"
)

var log = logger.Logger("config")

type Config struct {
	Port  int        `yaml:"port"`
	Redis conf.Redis `yaml:"redis"`
	Mysql conf.Mysql `yaml:"mysql"`
	Oss   conf.Oss   `yaml:"oss"`
}

var (
	config = new(Config)
)

func init() {
	logger.SetLogLevel("*", "INFO")
	var configPath string
	if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
		configPath = constant.ConfigFile
		log.Infof("use default config path: %v", configPath)
	} else {
		configPath = configEnv
		log.Infof("use env config path %v\n", configPath)
	}

	v := viper.New()
	v.SetConfigFile(configPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Info("config file changed:", e.Name)

		err = utils.YamlToStruct(configPath, &config)
		if err != nil {
			return
		}
		log.Infof("config: %+v", config)
	})

	err = utils.YamlToStruct(configPath, &config)
	if err != nil {
		panic("init config err")
	}

	log.Infof("config: %+v", config)
}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
