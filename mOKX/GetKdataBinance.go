package mOKX

import (
	"strconv"
	"strings"

	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX/binance"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type BinanceCandleDataType [12]any // Okx 原始数据
func GetKdataBinance(opt GetKdataOpt) (resData []TypeKd) {
	resData = []TypeKd{}
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

	fetchData, err := binance.FetchBinancePublic(binance.FetchBinancePublicOpt{
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

	var listStr []BinanceCandleDataType
	jsoniter.Unmarshal(fetchData, &listStr)

	rList := []TypeKd{}
	for _, item := range listStr {
		TimeStr := mStr.ToStr(mJson.ToJson(item[0]))
		kdata := TypeKd{
			InstID:   opt.InstID,
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(TimeStr, "0")),
			TimeStr:  mTime.UnixFormat(TimeStr),
			O:        mStr.ToStr(item[1]),
			H:        mStr.ToStr(item[2]),
			L:        mStr.ToStr(item[3]),
			C:        mStr.ToStr(item[4]),
			Vol:      mStr.ToStr(item[5]),
		}
		new_Kdata := NewKD(kdata, rList)
		rList = append(rList, new_Kdata)
	}
	if len(rList) == Size {
		resData = rList
	}

	return
}
