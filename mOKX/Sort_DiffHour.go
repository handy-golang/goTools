package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
)

// 按照 收盘价 价排序
func Sort_DiffHour(data []TypeWholeTickerAnaly) []TypeWholeTickerAnaly {
	size := len(data)
	list := make([]TypeWholeTickerAnaly, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := mStr.ToStr(list[j+1].DiffHour)
			b := mStr.ToStr(list[j].DiffHour)
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	// 设置 Idx 并翻转
	listIDX := []TypeWholeTickerAnaly{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Kdata := list[i]
		Kdata.DiffHour = j + 1
		listIDX = append(listIDX, Kdata)
		j++
	}
	return listIDX
}
