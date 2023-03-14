package testCase

import (
	"github.com/EasyGolang/goTools/internal/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
)

func GetKdata() {
	InstID := "ETH-USDT"
	/* 	resData := mOKX.GetKdataOKX(mOKX.GetKdataOpt{
	   		InstID: InstID,
	   		Page:   0, // 3 页 以后 就没有成交量了
	   		After:  mTime.GetUnixInt64(),
	   		Bar:    "1h", // 必须为 大写
	   	})
	   	mFile.Write(config.Dir.JsonData+"/okx_"+InstID+".json", mJson.ToStr(resData))
	*/
	/*
		resData := mOKX.GetKdataBinance(mOKX.GetKdataOpt{
			InstID: InstID,
			Page:   0, // 3 页 以后 就没有成交量了
			After:  mTime.GetUnixInt64(),
			Bar:    "1h", // 必须为 大写
		})
		mFile.Write(config.Dir.JsonData+"/bnb_"+InstID+".json", mJson.ToStr(resData))
	*/

	resData := mOKX.GetKdata(mOKX.GetKdataOpt{
		InstID: InstID,
		Page:   0, // 3 页 以后 就没有成交量了
		After:  0,
		Bar:    "1h", // 必须为 大写
	})
	mFile.Write(config.Dir.JsonData+"/merge_"+InstID+".json", mJson.ToStr(resData))
}
