package mOKX

import (
	"strings"

	"github.com/EasyGolang/goTools/mTime"
)

type KdataBarType struct {
	Interval int64  // 每一条数据之间间隔的毫秒数
	Okx      string // Okx 的值
	Binance  string // Bnb 的值
}

var KdataBarOpt = map[string]KdataBarType{
	"1m": {
		Interval: mTime.UnixTimeInt64.Minute * 1,
		Okx:      "1m",
		Binance:  "1m",
	},
	"3m": {
		Interval: mTime.UnixTimeInt64.Minute * 3,
		Okx:      "3m",
		Binance:  "3m",
	},
	"5m": {
		Interval: mTime.UnixTimeInt64.Minute * 5,
		Okx:      "5m",
		Binance:  "5m",
	},
	"15m": {
		Interval: mTime.UnixTimeInt64.Minute * 15,
		Okx:      "15m",
		Binance:  "15m",
	},
	"30m": {
		Interval: mTime.UnixTimeInt64.Minute * 30,
		Okx:      "30m",
		Binance:  "30m",
	},
	"1h": {
		Interval: mTime.UnixTimeInt64.Hour * 1,
		Okx:      "1H",
		Binance:  "1h",
	},
	"2h": {
		Interval: mTime.UnixTimeInt64.Hour * 2,
		Okx:      "2H",
		Binance:  "2h",
	},
	"4h": {
		Interval: mTime.UnixTimeInt64.Hour * 4,
		Okx:      "4H",
		Binance:  "4h",
	},
}

func GetBarOpt(b string) (resData KdataBarType) {
	param := strings.ToLower(b)
	bar := KdataBarOpt[param]
	return bar
}
