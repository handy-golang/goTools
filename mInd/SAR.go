package mInd

import (
	"fmt"

	"github.com/EasyGolang/goTools/mOKX"
)

/*
AF : 加速因子



*/

func SAR(KDList []mOKX.TypeKd) {
	period := 4 // 周期

	_, cutList := fundFirstSar(KDList, period)

	// AF_start := "0.02" // 加速因子初始值
	// AF_step := "0.02"  // 加速因子的上升值
	// AF_Max := "0.2"    // 加速因子最大值

	// trend := 1     // 默认为上涨
	// AF := AF_start // AF 初始值

	preArr := make([]mOKX.TypeKd, len(cutList)) // 上一个周期
	copy(preArr, cutList)

	// 极值计算
	Max := GetEP_H(preArr).H // 上一个周期的最高值
	Min := GetEP_L(preArr).L // 上一个周期的最低值

	fmt.Println(cutList[0].Time, cutList[len(cutList)-1].Time, Max, Min)
}

func fundFirstSar(KDList []mOKX.TypeKd, period int) (list []mOKX.TypeKd, cutList []mOKX.TypeKd) {
	cut := period

	if len(KDList) < period {
		cut = len(KDList)
	}

	cutList = KDList[:cut]
	list = KDList[cut:]

	return
}
