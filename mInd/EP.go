package mInd

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

// 求一端时间的最高价与最低价

// 最低价中的最低价
func GetArr_L(c []mOKX.TypeKd) mOKX.TypeKd {
	list := sortArrKd_L(c) // 最低价排序
	first := list[0]
	return first // 最低价的第一个
}

// 最高价中的最高价
func GetArr_H(c []mOKX.TypeKd) mOKX.TypeKd {
	list := sortArrKd_H(c) // 最高价排序
	last := list[len(list)-1]
	return last // 最高价排序的最后一个
}

// 按照格子的中心点来排序
func GetArrEP_Center(c []mOKX.TypeKd) (min mOKX.TypeKd, max mOKX.TypeKd) {
	list := sortArrKd_Center(c) // 收盘价排序
	first := list[0]
	last := list[len(list)-1]

	min = first
	max = last

	return
}

func sortArrKd_H(c []mOKX.TypeKd) []mOKX.TypeKd {
	size := len(c)
	list := make([]mOKX.TypeKd, size)
	copy(list, c)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].H
			b := list[j].H
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return list
}

func sortArrKd_L(c []mOKX.TypeKd) []mOKX.TypeKd {
	size := len(c)
	list := make([]mOKX.TypeKd, size)
	copy(list, c)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].L
			b := list[j].L
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return list
}

func sortArrKd_Center(c []mOKX.TypeKd) []mOKX.TypeKd {
	size := len(c)
	list := make([]mOKX.TypeKd, size)
	copy(list, c)
	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].CBas
			b := list[j].CBas
			if mCount.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return list
}
