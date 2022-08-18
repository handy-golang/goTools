package mInd

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

// 获取某一段的最低值
func GetEP_L(KDList []mOKX.TypeKd) (resKD mOKX.TypeKd) {
	resKD = mOKX.TypeKd{}
	if len(KDList) < 1 {
		return
	}
	resKD = KDList[0]
	for _, KD := range KDList {
		resVal := resKD.L
		itemVal := KD.L
		// 如果当前的 L 小于 resKD 则 替换， 否则跳过
		if mCount.Le(itemVal, resVal) < 0 {
			resKD = KD
		}
	}
	return
}

// 获取某一段的最高值
func GetEP_H(KDList []mOKX.TypeKd) (resKD mOKX.TypeKd) {
	resKD = mOKX.TypeKd{}
	if len(KDList) < 1 {
		return
	}
	resKD = KDList[0]
	for _, KD := range KDList {
		resVal := resKD.H
		itemVal := KD.H
		// 如果当前的 L 小于 resKD 则 替换， 否则跳过
		if mCount.Le(itemVal, resVal) > 0 {
			resKD = KD
		}
	}
	return
}
