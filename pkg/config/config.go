package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func NewConfig(env string) *viper.Viper {
	envConf := os.Getenv("GO_ENV")
	if envConf == "" {
		envConf = env
	}
	fmt.Println("load conf file", envConf)

	return getConfig(env)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
