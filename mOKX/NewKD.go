package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
)

// 构造新的 Kdata
func NewKD(now TypeKd, list []TypeKd) (kdata TypeKd) {
	kdata = now
	if mCount.Le(now.C, "0") < 1 {
		return
	}

	kdata.Dir = mCount.Le(kdata.C, kdata.O)

	kdata.CBas = mCount.Average([]string{now.H, now.L, now.C})

	kdata.HLPer = mCount.RoseCent(now.H, now.L)

	U_shade, D_shade := GetKdShade(kdata)
	kdata.U_shade = U_shade
	kdata.D_shade = D_shade

	Pre := kdata
	if len(list) > 1 {
		Pre = list[len(list)-1]
	}
	kdata.RosePer = mCount.RoseCent(now.C, Pre.C)
	kdata.C_dir = mCount.Le(now.CBas, Pre.CBas)

	return
}

func GetKdShade(now TypeKd) (U_shade, D_shade string) {
	if now.Dir > 0 { // 上涨时
		// 最高 - 收盘价 = 上影线
		U_shade = mCount.RoseCent(now.H, now.C)
		// 最低 - 开盘价 = 下影线
		D_shade = mCount.RoseCent(now.O, now.L)
	} else { // 下跌时
		// 最高 - 开盘价 = 上影线
		U_shade = mCount.RoseCent(now.H, now.O)
		// 最低 - 收盘价 = 下影线
		D_shade = mCount.RoseCent(now.C, now.L)
	}
	return
}
