package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)
	global.Start()

	fmt.Println(" =========  START  ========= ")

	// testCase.ClockStart()

	// testCase.GetSPOT()

	// testCase.GetKdata("EOS-USDT")

	// testCase.OKXFetch()
	// testCase.OKXWss()

	// testCase.CountTest()
	// testCase.StrFuzzy()

	// testCase.FileTest()

	// testCase.YaSuoDir()
	fmt.Println(" =========   END   ========= ")
}
