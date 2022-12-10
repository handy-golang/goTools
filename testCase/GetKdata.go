package testCase

import (
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX"
)

func GetKdata() {
	InstID := "ETH-USDT"
	resData := mOKX.GetKdata(mOKX.GetKdataOpt{
		InstID: InstID,
		Page:   0,
		After:  1609430400000,
		Bar:    "1h", // 必须为 大写
	})
	mFile.Write(config.Dir.JsonData+"/okx_"+InstID+".json", string(resData))
}
