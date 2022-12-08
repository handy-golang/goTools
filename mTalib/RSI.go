package mTalib

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mTalib/talib"
)

func RSI(opt CListOpt) string {
	n := opt.Period
	cLen := len(opt.CList)

	if cLen < n+1 {
		return "0"
	}

	dotNum := mCount.GetDecimal(opt.CList[0]) // 计算小数点位数
	var floatList []float64
	for _, val := range opt.CList {
		valDot := mCount.GetDecimal(opt.CList[0]) // 计算当前的小数点位数
		if valDot-dotNum > 0 {                    // 如果当前小数点位数大于现存小数点位数，则替换
			dotNum = valDot
		}
		floatVal := mCount.ToFloat(val, dotNum)
		floatList = append(floatList, floatVal) // 将数值完整的转化

		fmt.Println(val, floatVal)
	}

	// 计算 RSI
	rsiArr := talib.Rsi(floatList, n)
	rsiFloat := rsiArr[cLen-1]

	// 保留精确度，并转为字符串
	rsiStr := fmt.Sprintf("%f", rsiFloat)
	rsiStr = mCount.CentRound(rsiStr, dotNum)

	return rsiStr
}
