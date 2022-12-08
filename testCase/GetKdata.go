package testCase

import (
	"strings"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

var KdataList []mOKX.TypeKd

func GetKdata(InstID string) []mOKX.TypeKd {
	GetInst()

	InstInfo := GetInstInfo(InstID)
	KdataList = []mOKX.TypeKd{}

	if InstInfo.InstID != InstID {
		return KdataList
	}

	Kdata_file := mStr.Join(config.Dir.JsonData, "/", InstID, ".json")

	timeNow := mTime.GetUnix()
	// timeDiff := mCount.Mul("15", "300")
	// Minute := mCount.Mul(mTime.UnixTime.Minute, timeDiff)
	// after := mCount.Sub(timeNow, Minute)
	after := timeNow

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/market/candles",
		Data: map[string]any{
			"instId": InstID,
			"bar":    "1H",
			"after":  after,
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
			TimeStr:  mTime.UnixFormat(item[0]),
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

var EMA_18List = []string{}

func StorageKdata(kdata mOKX.TypeKd) {
	new_Kdata := mOKX.NewKD(kdata, KdataList)
	KdataList = append(KdataList, new_Kdata)

	cList_int := []float64{}
	cList := []string{}

	for _, val := range KdataList {
		cList_int = append(cList_int, mCount.ToFloat(val.C, -1))
		cList = append(cList, (val.C))
	}

	// EMA_18_int := []float64{float64(0)}
	// if len(cList_int) > 18 {
	// 	EMA_18_int = mTalib.Ema(cList_int, 18)
	// }

	// EMA_18 := mMetrics.EMA(mMetrics.EmaOpt{
	// 	CList:     cList,
	// 	Cycle:     18,
	// 	Precision: kdata.TickSz,
	// })

	// EMA_18List = append(EMA_18List, EMA_18)

	// CAP_3 := mMetrics.CAP(mMetrics.EmaOpt{
	// 	CList:     EMA_18List,
	// 	Cycle:     3,
	// 	Precision: kdata.TickSz,
	// })

	// global.KdataLog.Println(new_Kdata.TimeStr, EMA_18, EMA_18_int[len(EMA_18_int)-1])
	// global.KdataLog.Println(EMA_18_int)
	// global.KdataLog.Println(new_Kdata.TimeStr, CAP_3)

	// global.KdataLog.Println(mJson.Format(new_Kdata))
}
