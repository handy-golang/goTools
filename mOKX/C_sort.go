package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
)

// 按照 收盘价 价排序
func C_sort(data []Kd) []Kd {
	size := len(data)
	list := make([]Kd, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].C
			b := list[j].C
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
	listIDX := []Kd{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Ticker := list[i]
		Ticker.C_idx = j + 1
		listIDX = append(listIDX, Ticker)
		j++
	}
	return listIDX
}
