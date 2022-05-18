package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

var Dir struct {
	Home           string // Home 根目录
	HunterTrading  string // Hunter项目组
	App            string // APP 根目录
	ProdProject    string // 生产目录
	Client         string // 客户端目录
	Log            string // 日志文件目录
	HunterServer   string // Hunter 服务管理目录
	HunterReleases string // hunter.net 发布目录
}

var File struct {
	ServerEnv    string // /root/server_env.yaml
	AppServerEnv string // ./server_env.yaml
	UserConfig   string // ./user_config.yaml
}

var Constant struct {
	StaticOrigin string // 静态服务
}

func PathInit() {
	Constant.StaticOrigin = "//file.mo7.cc/HunterUserFile"

	Dir.Home = mPath.HomePath()
	Dir.HunterTrading = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"HunterTrading",
	)
	Dir.App = mStr.Join(
		Dir.HunterTrading,
		mStr.ToStr(os.PathSeparator),
		"DataCenter.net",
	)

	Dir.ProdProject = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"ProdProject",
	)
	Dir.Client = mStr.Join(
		Dir.HunterTrading,
		mStr.ToStr(os.PathSeparator),
		"Client",
	)

	Dir.Log = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"logs",
	)

	Dir.HunterServer = mStr.Join(
		Dir.ProdProject,
		mStr.ToStr(os.PathSeparator),
		"HunterServer",
	)
	Dir.HunterReleases = mStr.Join(
		Dir.HunterTrading,
		mStr.ToStr(os.PathSeparator),
		"HunterRelease",
	)

	File.ServerEnv = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"server_env.yaml",
	)
	File.AppServerEnv = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"server_env.yaml",
	)
	File.UserConfig = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"user_config.yaml",
	)
}

func main() {
	fmt.Println(" =========  START  ========= ")
	PathInit()

	Str := mJson.JsonFormat(mJson.ToJson(Dir))

	fmt.Println(Str)

	fmt.Println(" =========   END   ========= ")
}
