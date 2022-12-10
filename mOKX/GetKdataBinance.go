package mOKX

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mOKX/binance"
	"github.com/EasyGolang/goTools/mTime"
)

func GetKdataBinance(opt GetKdataOpt) (resData []OkxCandleDataType) {
	resData = []OkxCandleDataType{}
	if len(opt.InstID) < 2 {
		return
	}
	// Symbol
	Symbol := strings.Replace(opt.InstID, "-USDT", "USDT", 1)

	BarObj := GetBarOpt(opt.Bar)
	if BarObj.Interval < mTime.UnixTimeInt64.Minute {
		return
	}

	Size := 100

	// 时间设置
	now := mTime.GetUnixInt64()
	after := mTime.GetUnixInt64()
	// 时间必须大于6年前
	if opt.After > now-mTime.UnixTimeInt64.Day*2190 {
		after = opt.After
	}
	// 处理分页
	if opt.Page > 0 {
		pastTime := int64(opt.Page) * BarObj.Interval * int64(Size) // 一页数据 =  100 * 时间间隔
		after = after - pastTime                                    // 减去过去的时间节点
	}

	fmt.Println(Symbol)
	fmt.Println(BarObj.Binance)
	fmt.Println(after)
	fmt.Println(Size)

	fetchData, err := binance.FetchBinanceKdata(binance.FetchBinanceKdataOpt{
		Path:   "/api/v3/klines",
		Method: "get",
		Data: map[string]any{
			"symbol":   Symbol,
			"interval": BarObj.Binance,
			"endTime":  strconv.FormatInt(after, 10),
			"limit":    Size,
		},
	})
	if err != nil {
		return
	}

	mFile.Write(config.Dir.JsonData+"/bnb_"+Symbol+".json", string((fetchData)))

	return
}
