package mInd

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

func SAR(KDList []mOKX.TypeKd) (Value string, Dir int) {
	// 周期为 4
	period := 4
	list, cutList := fundFirstSar(KDList, period)

	AF_init := "0.02" // 加速因子初始值
	AF_step := "0.02" // 加速因子的上升值
	AF_Max := "0.2"   // 加速因子最大值

	trend := 1 // 默认为上涨
	AF := "0.02"

	preArr := make([]mOKX.TypeKd, len(cutList)) // 上一个周期
	copy(preArr, cutList)

	// 极值计算
	Max := GetArr_H(preArr).H // 上一个周期的最高值
	Min := GetArr_L(preArr).L // 上一个周期的最低值

	SarVal := Min // 初始的SAR值为最低点
	EP := Max

	// 准备工具函数

	// 更新加速因子 并 限制最大值
	AFUpdate := func() {
		AF = mCount.Add(AF, AF_step)
		if mCount.Le(AF, AF_Max) > 0 {
			AF = AF_Max
		}
	}

	nowArr := []mOKX.TypeKd{}

	var precision int32 // 精度
	for _, val := range list {
		precision = mCount.GetDecimal(val.TickSz)

		nowArr = append(nowArr, val)

		if trend > 0 { // 上升计算
			if mCount.Le(SarVal, val.L) > 0 {
				// 翻转
				trend = -1

				AF = AF_init // 重置加速因子

				preArr = make([]mOKX.TypeKd, len(nowArr)) // 记录当前周期的值
				copy(preArr, nowArr)
				nowArr = []mOKX.TypeKd{} // 重置当前的周期值

				Max = GetArr_H(preArr).H // 上一个周期的最高值
				// Min = GetArr_L(periodArr).L // 上一个周期的最低值
				Min = val.L // 上一个周期的最低值

				SarVal = Max // 上个周期的最高价格  这里正确了
				EP = Min

				nowArr = append(nowArr, val) // 记录当前的周期值
				continue

			} else {
				// 没有翻转 , 出现新高，
				if mCount.Le(val.H, Max) > 0 {
					Max = val.H // 更新最高价为当前最高价
					EP = Max

					AFUpdate() // 更新加速因子
				}
			}
		} else { // 下跌算法
			if mCount.Le(SarVal, val.H) < 0 {
				// 翻转
				trend = 1

				AF = AF_init // 重置加速因子

				preArr = make([]mOKX.TypeKd, len(nowArr)) // 记录当前周期的值
				copy(preArr, nowArr)

				nowArr = []mOKX.TypeKd{} // 重置当前的周期值

				// 清除数组
				Min = GetArr_L(preArr).L // 上一个周期的最低值
				// Max = GetArr_H(periodArr).H // 上一个周期的最高值
				Max = val.H // 上一个周期的最高值

				SarVal = Min // 上个周期的最低价 这里正确了
				EP = Max

				nowArr = append(nowArr, val) // 记录当前的周期值
				continue

			} else {
				if mCount.Le(val.L, Min) < 0 {
					Min = val.L // 更新 最低价格为当前最低价格
					EP = Min

					AFUpdate() // 更新加速因子
				}
			}
		}

		SarVal = SarCount(SarVal, AF, EP)
		SarVal = mCount.CentRound(SarVal, precision) // 保留精确度
	}

	Value = SarVal
	Dir = trend

	return "1", 0
}

func SarCount(sar, af, ep string) string {
	a := mCount.Sub(ep, sar) // ep - sar

	b := mCount.Mul(af, a) // af * (ep - sar)

	c := mCount.Add(sar, b) // sar + af * (ep - sar)

	return c
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
