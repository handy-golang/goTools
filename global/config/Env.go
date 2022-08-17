package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Port    int    `json:"Port"`
}

var AppEnv struct {
	RunMod int
}

func LoadAppEnv() {
	viper.SetConfigFile(File.AppEnv)

	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Errorf("AppEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&AppEnv)
}
