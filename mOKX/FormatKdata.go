package mOKX

import (
	"fmt"

	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type FormatKdataParam struct {
	Data     any    // [][9]string
	InstID   string // 产品信息
	DataInfo string // 格式化后的描述
}

func FormatKdata(opt FormatKdataParam) []TypeKd {
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
