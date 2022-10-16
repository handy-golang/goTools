package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
)

// 涨跌幅排序
func Sort_RosePer(data []AnalySliceType) []AnalySliceType {
	size := len(data)
	list := make([]AnalySliceType, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].RosePer
			b := list[j].RosePer
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	// 设置 U_RIdx 并翻转
	listIDX := []AnalySliceType{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Analy := list[i]
		listIDX = append(listIDX, Analy)
		j++
	}
	return listIDX
}
