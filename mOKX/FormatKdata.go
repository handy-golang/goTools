package mOKX

import (
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type FormatKdataParam struct {
	Data     any      // [][8]string
	Inst     TypeInst // 产品信息
	DataType string   // 格式化后的描述
}

func FormatKdata(opt FormatKdataParam) []TypeKd {
	KdataList := []TypeKd{} // 声明存储
	// 检查参数
	errStr := []string{}
	switch {
	case len(opt.Inst.InstID) < 2:
		errStr = append(errStr, "Inst")
		fallthrough
	case len(opt.DataType) < 2:
		errStr = append(errStr, "DataType")
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

	CcyName := opt.Inst.InstID
	if opt.Inst.InstType == "SWAP" {
		CcyName = strings.Replace(opt.Inst.InstID, config.SWAP_suffix, "", -1)
	}
	if opt.Inst.InstType == "SPOT" {
		CcyName = strings.Replace(opt.Inst.InstID, config.SPOT_suffix, "", -1)
	}

	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]
		kdata := TypeKd{
			InstID:   opt.Inst.InstID,
			CcyName:  CcyName,
			TickSz:   opt.Inst.TickSz,
			InstType: opt.Inst.InstType,
			CtVal:    opt.Inst.CtVal,
			MinSz:    opt.Inst.MinSz,
			MaxMktSz: opt.Inst.MaxMktSz,
			TimeStr:  mTime.UnixFormat(item[0]),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(item[0], "0")),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
			Vol:      item[5],
			DataType: opt.DataType,
		}
		new_Kdata := NewKD(kdata, KdataList)
		KdataList = append(KdataList, new_Kdata)
	}

	return KdataList
}
