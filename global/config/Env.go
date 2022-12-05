package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var AppInfo struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
	Port    int    `bson:"Port"`
}

var AppEnv struct {
	RunMod        int
	MongoAddress  string
	MongoPassword string
	MongoUserName string
}

var BinanceKey = struct {
	ApiKey    string `bson:"ApiKey"`
	SecretKey string `bson:"SecretKey"`
}{
	ApiKey:    "e1ch7VN2WEml",
	SecretKey: "E9w2QmLtV9e",
}

func LoadAppEnv() {
	viper.SetConfigFile(File.AppEnv)

	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Errorf("AppEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
	}
	viper.Unmarshal(&AppEnv)
}
