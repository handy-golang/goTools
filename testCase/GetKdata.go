package testCase

import (
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mInd"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

var KdataList []mOKX.TypeKd

func GetKdata(InstID string) []mOKX.TypeKd {
	InstInfo := GetInstInfo(InstID)
	KdataList = []mOKX.TypeKd{}

	if InstInfo.InstID != InstID {
		return KdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/candles",
		Data: map[string]any{
			"instId": InstID,
			"bar":    "15m",
			"after":  mTime.GetUnix(),
			"limit":  300,
		},
		Method:        "get",
		LocalJsonPath: Kdata_file,
		IsLocalJson:   config.AppEnv.RunMod == 1,
	})
	if err != nil {
		global.LogErr("kdata.GetKdata", InstID, err)
		return nil
	}
	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("kdata.GetKdata", InstID, "Err", result)
		return nil
	}

	FormatKdata(result.Data, InstInfo)

	if len(KdataList) != 300 {
		global.LogErr("kdata.GetKdata resData", mStr.ToStr(resData))
	}

	// 写入数据文件
	mFile.Write(Kdata_file, mStr.ToStr(resData))
	return KdataList
}

func FormatKdata(data any, Inst mOKX.TypeInst) {
	var list []mOKX.CandleDataType
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	global.LogErr("kdata.FormatKdata", len(list), Inst.InstID)

	CcyName := Inst.InstID
	if Inst.InstType == "SWAP" {
		CcyName = strings.Replace(Inst.InstID, config.SWAP_suffix, "", -1)
	}
	if Inst.InstType == "SPOT" {
		CcyName = strings.Replace(Inst.InstID, config.SPOT_suffix, "", -1)
	}

	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]

		kdata := mOKX.TypeKd{
			InstID:   Inst.InstID,
			CcyName:  CcyName,
			TickSz:   Inst.TickSz,
			InstType: Inst.InstType,
			CtVal:    Inst.CtVal,
			MinSz:    Inst.MinSz,
			MaxMktSz: Inst.MaxMktSz,
			Time:     mTime.MsToTime(item[0], "0"),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(item[0], "0")),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
			Vol:      item[5],
			VolCcy:   item[6],
			DataType: "GetKdata",
		}
		StorageKdata(kdata)
	}
}

func StorageKdata(kdata mOKX.TypeKd) {
	new_Kdata := mOKX.NewKD(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)

	Value, Dir := mInd.SAR(KdataList)

	fmt.Println(new_Kdata.InstID, new_Kdata.Time, new_Kdata.C, Value, Dir)

	global.KdataLog.Println(mJson.Format(new_Kdata))
}
