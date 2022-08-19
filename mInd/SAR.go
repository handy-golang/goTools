package mInd

import (
	"github.com/EasyGolang/goTools/mOKX"
)

/*
https://baike.baidu.com/item/SAR%E6%8C%87%E6%A0%87/6329095

*/

var (
	Period   = 4      // 周期
	AF_start = "0.02" // 加速因子初始值
	AF_step  = "0.02" // 加速因子的上升值
	AF_Max   = "0.2"  // 加速因子最大值
)

func SAR(KDList []mOKX.TypeKd) (SarVal string, trend int) {
	SarVal = "0"
	trend = 1

	return
}
