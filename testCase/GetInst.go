package testCase

import (
	"strings"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

var (
	SPOT_list = make(map[string]mOKX.TypeInst)
	SWAP_list = make(map[string]mOKX.TypeInst)
)

func GetInst() {
	GetSPOT()
	GetSWAP()
}

func GetInstInfo(InstID string) (resData mOKX.TypeInst) {
	resData = mOKX.TypeInst{}

	for _, item := range SPOT_list {
		if item.InstID == InstID {
			resData = item
		}
	}

	for _, item := range SWAP_list {
		if item.InstID == InstID {
			resData = item
		}
	}

	return resData
}

func GetSPOT() {
	SPOT_file := mStr.Join(config.Dir.JsonData, "/SPOT.json")

	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/public/instruments",
		Data: map[string]any{
			"instType": "SPOT",
		},
		Method:        "get",
		LocalJsonPath: SPOT_file,
		IsLocalJson:   config.AppEnv.RunMod == 1,
	})
	if err != nil {
		global.LogErr("SPOT", err)
		return
	}
	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("SPOT-err", result)
		return
	}

	setSPOT_list(result.Data)

	// 写入数据文件
	mFile.Write(SPOT_file, mStr.ToStr(resData))
}

func setSPOT_list(data any) {
	var list []mOKX.TypeInst
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	global.KdataLog.Println("inst.setSPOT_list", len(list), "SPOT")

	for _, val := range list {
		find := strings.Contains(val.InstID, config.SPOT_suffix)
		if find && val.State == "live" {
			SPOT_list[val.InstID] = val
		}
	}
}

// 获取可交易合约列表
func GetSWAP() {
	SWAP_file := mStr.Join(config.Dir.JsonData, "/SWAP.json")
	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/public/instruments",
		Data: map[string]any{
			"instType": "SWAP",
		},
		Method:        "get",
		LocalJsonPath: SWAP_file,
		IsLocalJson:   true,
	})
	if err != nil {
		global.LogErr("inst.SWAP ", err)
		return
	}

	var result mOKX.TypeReq
	jsoniter.Unmarshal(resData, &result)
	if result.Code != "0" {
		global.LogErr("SPOT-err", result)
		return
	}

	setSWAP_list(result.Data)

	// 写入日志文件
	mFile.Write(SWAP_file, mStr.ToStr(resData))
}

func setSWAP_list(data any) {
	var list []mOKX.TypeInst
	jsonStr := mJson.ToJson(data)
	jsoniter.Unmarshal(jsonStr, &list)

	global.KdataLog.Println("inst.setSWAP_list", len(list), "SWAP")

	for _, val := range list {

		find := strings.Contains(val.InstID, config.SWAP_suffix) // 统一计价单位
		if find && val.State == "live" {
			SWAP_list[val.InstID] = val
		}
	}
}
