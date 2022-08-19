package mInd

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

/*
AF : 加速因子



*/

var (
	Period   = 4      // 周期
	AF_start = "0.02" // 加速因子初始值
	AF_step  = "0.02" // 加速因子的上升值
	AF_Max   = "0.2"  // 加速因子最大值
)

// 废弃，不能用， SAR 指标太难了
func SAR(KDList []mOKX.TypeKd) (SarVal string, trend int) {
	list, cutList := fundFirstSar(KDList) // 数组切片

	// 极值计算
	preArr := make([]mOKX.TypeKd, len(cutList)) // 上一个周期 的数组
	copy(preArr, cutList)
	Max := GetEP_H(preArr).H // 上一个周期的最高值
	Min := GetEP_L(preArr).L // 上一个周期的最低值

	SarVal = Min   // 初始的 SAR 值为 上一个周期最低点
	trend = 1      // 默认为上涨
	AF := AF_start // AF 初始值
	EP := Max      // 初始 EP 值为 上一个周期最高点

	// 工具函数 : 更新加速因子 并 限制最大值
	AFUpdate := func() {
		AF = mCount.Add(AF, AF_step)
		if mCount.Le(AF, AF_Max) > 0 {
			AF = AF_Max
		}
	}

	nowArr := []mOKX.TypeKd{} // 当前正在遍历的Arr

	var precision int32
	for _, item := range list {
		precision = mCount.GetDecimal(item.TickSz)
		nowArr = append(nowArr, item)
		Max, Min := getPreMax(nowArr)

		if trend > 0 { // 上升计算
			/*
			  SarVal = SarVal + AF *  (EP - SarVal)
			  SAR(t+1)=SAR(t)+Af(t)*(Ep(t) – SAR(t))
			*/
			EP_Sub_Sar := mCount.Sub(EP, SarVal)
			AF_Mul_ESS := mCount.Mul(AF, EP_Sub_Sar)
			SarVal = mCount.Add(SarVal, AF_Mul_ESS)
			SarVal = mCount.CentRound(SarVal, precision)

			if mCount.Le(SarVal, item.L) > 0 {
				trend = -1    // 翻转为跌势
				AF = AF_start // AF 初始值
				SarVal = Max
				EP = Min

			} else {
				trend = 1 // 涨势
				EP = Max
				if mCount.Le(item.H, Max) > 0 {
					AFUpdate() // 更新加速因子
				}

			}
		} else { // 下跌计算

			/*
			  SarVal = SarVal + AF *  (EP - SarVal)
			  SAR(t+1)=SAR(t)+Af(t)*(Ep(t) – SAR(t))
			*/
			EP_Sub_Sar := mCount.Sub(EP, SarVal)
			AF_Mul_ESS := mCount.Mul(AF, EP_Sub_Sar)
			SarVal = mCount.Add(SarVal, AF_Mul_ESS)
			SarVal = mCount.CentRound(SarVal, precision)
			if mCount.Le(SarVal, item.H) < 0 {
				trend = 1     // 翻转为涨势
				AF = AF_start // AF 初始值
				SarVal = Min
				EP = Max

			} else {
				trend = -1 // 继续跌势
				EP = Min
				if mCount.Le(item.L, Min) < 0 {
					AFUpdate() // 更新加速因子
				}

			}
		}

	}

	return
}

func fundFirstSar(KDList []mOKX.TypeKd) (list []mOKX.TypeKd, cutList []mOKX.TypeKd) {
	cut := Period

	if len(KDList) < Period {
		cut = len(KDList)
	}

	cutList = KDList[:cut]
	list = KDList[cut:]

	return
}

func getPreMax(KDList []mOKX.TypeKd) (Max, Min string) {
	Len := len(KDList)

	if Len < Period {
		Len = Period
	}

	cutList := KDList[Len-Period:]

	Max = GetEP_H(cutList).H // 上一个周期的最高值
	Min = GetEP_L(cutList).L // 上一个周期的最低值

	return
}
