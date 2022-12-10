package mOKX

import (
	"fmt"
	"strconv"

	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type GetKdataOpt struct {
	InstID string `bson:"InstID"`
	After  int64  `bson:"After"` // 此时间之前的内容
	Page   int    `bson:"Page"`  // 往前第几页
	Bar    string `bson:"Bar"`   // 1m/3m/5m/15m/30m/1h/2h/4h
}

func GetKdataOKX(opt GetKdataOpt) (resData []OkxCandleDataType) {
	resData = []OkxCandleDataType{}
	if len(opt.InstID) < 2 {
		return
	}

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

	// 判断应该采取哪个接口获取数据  after 距离 now 有多少条数据?
	path := "/api/v5/market/candles"
	fromNowItem := (now - after) / BarObj.Interval
	if fromNowItem > 300 {
		path = "/api/v5/market/history-index-candles"
	}

	fetchData, err := FetchOKX(OptFetchOKX{
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
	jsoniter.Unmarshal(fetchData, &result)
	if result.Code != "0" {
		return
	}

	jsonByte := mJson.ToJson(result.Data)

	var list []OkxCandleDataType
	err = jsoniter.Unmarshal(jsonByte, &list)
	if err != nil {
		return
	}
	if len(list) != Size {
		return
	}

	resData = list

	return
}

type FormatOkxKdataParam struct {
	Data     any    // [][9]string
	InstID   string // 产品信息
	DataInfo string // 格式化后的描述
}

func FormatOkxKdata(opt FormatOkxKdataParam) []TypeKd {
	KdataList := []TypeKd{} // 声明存储
	// 检查参数
	errStr := []string{}
	switch {
	case len(opt.InstID) < 2:
		errStr = append(errStr, "Inst")
		fallthrough
	case len(opt.DataInfo) < 2:
		errStr = append(errStr, "DataInfo")
	case len(mStr.ToStr(opt.Data)) < 30:
		errStr = append(errStr, "Data")
	}
	if len(errStr) > 0 {
		errStr := fmt.Errorf("缺少参数:%+v", errStr)
		panic(errStr)
	}

	// 解析 List
	var list []OkxCandleDataType
	jsonStr := mJson.ToJson(opt.Data)
	err := jsoniter.Unmarshal(jsonStr, &list)
	if err != nil {
		return KdataList
	}

	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]
		kdata := TypeKd{
			InstID:   opt.InstID,
			TimeStr:  mTime.UnixFormat(item[0]),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(item[0], "0")),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
			Vol:      item[5],
			DataInfo: opt.DataInfo,
		}
		new_Kdata := NewKD(kdata, KdataList)
		KdataList = append(KdataList, new_Kdata)
	}

	return KdataList
}
