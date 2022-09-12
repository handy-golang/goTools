package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
)

func FileTest() {
	resData, err := mFile.Tinypng(mFile.TinyOpt{
		Src:    config.Dir.App + "/jsonData/2.png",
		Email:  "meichangliang@mo7.cc",
		ApiKey: "Hl6wpxNdBg0Dvv2s7BcVsKks1tFZ2wBl",
	})

	fmt.Println("resData", resData)
	fmt.Println("err", err)
}
