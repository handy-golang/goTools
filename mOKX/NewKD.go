package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
)

// 构造新的 Kdata
func NewKD(now TypeKd, list []TypeKd) (kdata TypeKd) {
	kdata = now

	if mCount.Le("0", now.C) > -1 {
		return
	}

	kdata.Dir = mCount.Le(kdata.C, kdata.O)

	kdata.CBas = mCount.Average([]string{now.H, now.L, now.C})

	kdata.HLPer = mCount.RoseCent(now.H, now.L)

	U_shade, D_shade := NewKdShade(kdata)
	kdata.U_shade = U_shade
	kdata.D_shade = D_shade

	if len(list) < 1 {
		return
	}
	Pre := list[len(list)-1]
	kdata.RosePer = mCount.RoseCent(now.C, Pre.C)
	kdata.C_dir = NewKddC_dir(kdata, Pre)

	// 复制数组
	size := len(list)
	newList := make([]TypeKd, size)
	copy(newList, list)
	newList = append(newList, kdata)

	return
}

func NewKdShade(now TypeKd) (U_shade, D_shade string) {
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

func NewKddC_dir(now, pre TypeKd) int {
	// 格子方向
	C_dir := mCount.Le(now.CBas, pre.CBas) // 以中心点为基准来计算，当前-过去的
	return C_dir
}
