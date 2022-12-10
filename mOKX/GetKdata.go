package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mTime"
)

type GetKdataOpt struct {
	InstID string `bson:"InstID"`
	After  int64  `bson:"After"` // 此时间之前的内容
	Page   int    `bson:"Page"`  // 往前第几页
	Bar    string `bson:"Bar"`   // 1m/3m/5m/15m/30m/1h/2h/4h
}

func GetKdata(opt GetKdataOpt) (KdataList []TypeKd) {
	KdataList = []TypeKd{}
	if len(opt.InstID) < 2 {
		return
	}
	BarObj := GetBarOpt(opt.Bar)
	if BarObj.Interval < mTime.UnixTimeInt64.Minute {
		return
	}

	BinanceList := GetKdataBinance(opt)
	OKXList := GetKdataOKX(opt)

	for _, item := range OKXList {
		OkxItem := item
		for _, BinanceItem := range BinanceList {
			if OkxItem.TimeUnix == BinanceItem.TimeUnix {
				Vol := mCount.Add(OkxItem.Vol, BinanceItem.Vol)
				OkxItem.Vol = Vol
				break
			}
		}

		KdataList = append(KdataList, OkxItem)
	}

	return
}
