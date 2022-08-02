package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)
	HelloWord()
	global.Start()

	fmt.Println(" =========  START  ========= ")

	// testCase.OKXFetch()

	fmt.Println(" =========   END   ========= ")
}

func HelloWord() {
	Tmpl := ` 
	#########################################################
	####                欢迎使用 goTools                   ###     
	####       github.com/EasyGolang/goTools v${version}      ###     
	####              go clean --modcache                 ###     
	####          作者邮箱 meichangliang@mo7.cc            ###      
	####  项目地址 https://github.com/EasyGolang/goTools   ###     
	#########################################################
	`
	lMap := map[string]string{
		"version": config.AppInfo.Version,
	}

	logStr := mStr.Temp(Tmpl, lMap)

	fmt.Println(logStr)
}
