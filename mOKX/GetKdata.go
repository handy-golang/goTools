package mOKX

import (
	"strconv"

	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type GetKdataOpt struct {
	InstID string `bson:"InstID"`
	After  int64  `bson:"After"` // 此时间之前的内容
	Page   int    `bson:"Page"`  // 往前第几页
	Bar    string `bson:"Bar"`   // 1m/3m/5m/15m/30m/1h/2h/4h
}

func GetKdata(opt GetKdataOpt) {
	if len(opt.InstID) < 2 {
		return
	}

	BarObj := GetBarOpt(opt.Bar)
	if BarObj.Interval < mTime.UnixTimeInt64.Minute {
		return
	}

	Size := 100

	// 时间设置
	after := mTime.GetUnixInt64()
	if opt.After > 946656000000 {
		after = opt.After
	}

	// 处理分页
	if opt.Page > 0 {
	}

	// 判断应该采取哪个接口获取数据
	path := "/api/v5/market/candles"

	// fmt.Println(after)

	// path = "/api/v5/market/history-index-candles"

	resData, err := FetchOKX(OptFetchOKX{
		Path: path,
		Data: map[string]any{
			"instId": opt.InstID,
			"bar":    BarObj.Okx,
			"after":  strconv.FormatInt(after, 10),
			"limit":  Size,
		},
		Method: "get",
	})
	if err != nil {
		return
	}
	var result TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		return
	}

	mFile.Write(config.Dir.JsonData+"/okx_"+opt.InstID+".json", string(mJson.ToJson(result.Data)))

	return
}
