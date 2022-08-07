package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
)

// 构造新的 Kdata
func AnalyNewKd(now TypeKd, list []TypeKd) (kdata TypeKd) {
	kdata = TypeKd{
		InstID:   now.InstID,
		CcyName:  now.CcyName,
		TimeUnix: now.TimeUnix,
		Time:     now.Time,
		O:        now.O,
		H:        now.H,
		L:        now.L,
		C:        now.C,
		VolCcy:   now.VolCcy,
		Type:     now.Type,
	}

	if mCount.Le("0", now.C) > -1 {
		return
	}

	kdata.Dir = mCount.Le(kdata.C, kdata.O)

	Center := mCount.Average([]string{now.C, now.O, now.H, now.L})
	kdata.Center = mCount.PriceCent(Center, now.C)

	kdata.HLPer = mCount.RoseCent(now.H, now.L)

	U_shade, D_shade := NewKdShade(kdata)
	kdata.U_shade = mCount.PriceCent(U_shade, now.C)
	kdata.D_shade = mCount.PriceCent(D_shade, now.C)

	if len(list) < 1 {
		return
	}
	Pre := list[len(list)-1]
	kdata.RosePer = mCount.RoseCent(now.C, Pre.C)
	kdata.C_dir = NewKddC_dir(kdata, Pre)

	return
}

func NewKdShade(now TypeKd) (U_shade, D_shade string) {
	if now.Dir > 0 { // 上涨时
		// 最高 - 收盘价 = 上影线
		U_shade = mCount.Rose(now.H, now.C)
		// 最低 - 开盘价 = 下影线
		D_shade = mCount.Rose(now.O, now.L)
	} else { // 下跌时
		// 最高 - 开盘价 = 上影线
		U_shade = mCount.Rose(now.H, now.O)
		// 最低 - 收盘价 = 下影线
		D_shade = mCount.Rose(now.C, now.L)
	}
	return
}

func NewKddC_dir(now, pre TypeKd) int {
	// 格子方向
	C_dir := mCount.Le(now.Center, pre.Center) // 以中心点为基准来计算，当前-过去的
	return C_dir
}
