package mInd

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

/*
AF : 加速因子



*/

func SAR(KDList []mOKX.TypeKd) (SarVal string, trend int) {
	period := 4 // 周期

	list, cutList := fundFirstSar(KDList, period) // 数组切片

	AF_start := "0.02" // 加速因子初始值
	AF_step := "0.02"  // 加速因子的上升值
	AF_Max := "0.2"    // 加速因子最大值

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

	for _, item := range list {
		nowArr = append(nowArr, item)

		if trend > 0 { // 上升计算
			/*
			  SarVal = SarVal + AF *  (EP - SarVal)
			  SAR(t+1)=SAR(t)+Af(t)*(Ep(t) – SAR(t))
			*/
			EP_Sub_Sar := mCount.Sub(EP, SarVal)
			AF_Mul_ESS := mCount.Mul(AF, EP_Sub_Sar)
			SarVal = mCount.Add(SarVal, AF_Mul_ESS)

			if mCount.Le(SarVal, item.L) > 0 {
				trend = -1 // 翻转为跌势

				Max = GetEP_H(nowArr).H // 上一个周期的最高值
				Min = GetEP_L(nowArr).L // 上一个周期的最低值

				SarVal = Max  // 前段时间的最高价格
				AF = AF_start // AF 初始值
				EP = Min
				nowArr = []mOKX.TypeKd{} // 清空

			} else {
				trend = 1 // 涨势
				EP = Max
				if mCount.Le(item.H, Max) > 0 {
					AFUpdate() // 更新加速因子
				}
			}
			continue
		}

		if trend < 0 { // 下跌计算
			EP_Sub_Sar := mCount.Sub(EP, SarVal)
			AF_Mul_ESS := mCount.Mul(AF, EP_Sub_Sar)
			SarVal = mCount.Add(SarVal, AF_Mul_ESS)

			if mCount.Le(SarVal, item.H) < 0 {
				trend = 1               // 翻转为涨势
				Max = GetEP_H(nowArr).H // 上一个周期的最高值
				Min = GetEP_L(nowArr).L // 上一个周期的最低值

				SarVal = Min  // 前段时间的最高价格
				AF = AF_start // AF 初始值
				EP = Max
				nowArr = []mOKX.TypeKd{} // 清空

			} else {
				trend = -1 // 继续跌势
				EP = Min
				if mCount.Le(item.L, Min) < 0 {
					AFUpdate() // 更新加速因子
				}
			}

			continue
		}
	}

	return
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
