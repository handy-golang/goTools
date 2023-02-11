package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/testCase"
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

	// testCase.GetKdata()

	// testCase.OKXFetch()
	// testCase.OKXWss()

	// testCase.CountTest()
	// testCase.StrFuzzy()

	// testCase.FileTest()

	// testCase.YaSuoDir()
	// OrganizeDatabase()

	// testCase.StartDBRun()

	// testCase.IPtest()

	// testCase.BalanceFetch()

	// testCase.TimeTest()

	// testCase.Tactics()

	// testCase.ShellTest()

	testCase.TestEmail()

	// testCase.TaskTest()

	fmt.Println(" =========   END   ========= ")
}
