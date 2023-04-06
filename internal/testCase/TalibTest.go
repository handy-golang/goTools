package testCase

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/internal/global"
	"github.com/EasyGolang/goTools/internal/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mTalib"
	"github.com/EasyGolang/goTools/mTime"
)

func TalibTest() {
	// 获取数据
	InstID := "ETH-USDT"
	Page := 5

	KdataList := []mOKX.TypeKd{}
	for i := Page; i >= 0; i-- {
		time.Sleep(time.Second / 3)
		List := mOKX.GetKdata(mOKX.GetKdataOpt{
			InstID: InstID,
			Page:   i,
			After:  mTime.GetUnixInt64(),
		})
		KdataList = append(KdataList, List...)
	}

	for key, val := range KdataList {
		preIndex := key - 1
		if preIndex < 0 {
			preIndex = 0
		}
		preItem := KdataList[preIndex]
		nowItem := KdataList[key]
		if key > 0 {
			if nowItem.TimeUnix-preItem.TimeUnix != mTime.UnixTimeInt64.Hour {
				fmt.Println("数据检查出错", val.InstID, val.TimeStr, key)
				break
			}
		}
	}

	mFile.Write(config.Dir.JsonData+"/merge_"+InstID+".json", mJson.ToStr(KdataList))

	countKdata := []mOKX.TypeKd{}
	for _, item := range KdataList {
		countKdata = append(countKdata, item)
		Ema := mTalib.ClistNew(mTalib.ClistOpt{
			KDList: countKdata,
			Period: 171,
		}).EMA().ToStr()

		Ma := mTalib.ClistNew(mTalib.ClistOpt{
			KDList: countKdata,
			Period: 171,
		}).MA().ToStr()

		global.Log.Println(item.TimeStr, item.C, "EMA:", Ema, "MA", Ma)
	}
}
