package mCount

import "fmt"

// 多个数之间求和
func Sum(n []string) string {
	s := "0"
	for _, v := range n {
		s = Add(s, v)
	}
	return s
}

// 算术平均数
func Average(sl []string) string {
	all := Sum(sl)
	n := len(sl)
	r := Div(all, fmt.Sprint(n))
	return r
}
