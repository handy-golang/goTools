package mOKX

import (
	"strings"

	"github.com/EasyGolang/goTools/mTime"
)

type KdataBarType struct {
	Interval int64  // 每一条数据之间间隔的毫秒数
	Okx      string // Okx 的值
	Bnb      string // Bnb 的值
}

var KdataBarOpt = map[string]KdataBarType{
	"1m": {
		Interval: mTime.UnixTimeInt64.Minute * 1,
		Okx:      "1m",
		Bnb:      "1m",
	},
	"3m": {
		Interval: mTime.UnixTimeInt64.Minute * 3,
		Okx:      "3m",
		Bnb:      "3m",
	},
	"5m": {
		Interval: mTime.UnixTimeInt64.Minute * 5,
		Okx:      "5m",
		Bnb:      "5m",
	},
	"15m": {
		Interval: mTime.UnixTimeInt64.Minute * 15,
		Okx:      "15m",
		Bnb:      "15m",
	},
	"30m": {
		Interval: mTime.UnixTimeInt64.Minute * 30,
		Okx:      "30m",
		Bnb:      "30m",
	},
	"1h": {
		Interval: mTime.UnixTimeInt64.Hour * 1,
		Okx:      "1H",
		Bnb:      "1h",
	},
	"2h": {
		Interval: mTime.UnixTimeInt64.Hour * 2,
		Okx:      "2H",
		Bnb:      "2h",
	},
	"4h": {
		Interval: mTime.UnixTimeInt64.Hour * 4,
		Okx:      "4H",
		Bnb:      "4h",
	},
}

func GetBarOpt(b string) (resData KdataBarType) {
	param := strings.ToLower(b)
	bar := KdataBarOpt[param]
	return bar
}
