package testCase

import (
	"github.com/EasyGolang/goTools/mOKX"
)

func GetKdata(InstID string) {
	mOKX.GetKdata(mOKX.GetKdataOpt{
		InstID: "ETH-USDT",
		Page:   0,
		After:  0,
		Bar:    "1m", // 必须为 大写
	})
}
