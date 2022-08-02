package mOKX

import (
	"github.com/EasyGolang/goTools/mCount"
)

// 涨跌幅排序
func U_R24Sort(data []TickerType) []TickerType {
	size := len(data)
	list := make([]TickerType, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].U_R24
			b := list[j].U_R24
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
	listIDX := []TickerType{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Ticker := list[i]
		Ticker.U_RIdx = j + 1
		listIDX = append(listIDX, Ticker)
		j++
	}
	return listIDX
}
