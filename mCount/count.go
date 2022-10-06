package mCount

import (
	"strings"

	"github.com/shopspring/decimal"
)

func toDec(s string) decimal.Decimal {
	n, _ := decimal.NewFromString(s)
	return n
}

// a+b
func Add(a, b string) string {
	n := toDec(a).Add(toDec(b))
	return n.String()
}

// a-b
func Sub(a, b string) string {
	n := toDec(a).Sub(toDec(b))
	return n.String()
}

// a*b
func Mul(a, b string) string {
	n := toDec(a).Mul(toDec(b))
	return n.String()
}

// a/b
func Div(a, b string) string {
	if Le(b, "0") == 0 {
		return b
	}
	n := toDec(a).Div(toDec(b))
	return n.String()
}

// (a/b)*100
func Per(a, b string) string {
	if Le(b, "0") == 0 {
		return b
	}
	n := toDec(a).Div(toDec(b)).Mul(toDec("100"))
	return n.String()
}

// (a/b)*100 保留三位小数
func PerCent(a, b string) string {
	if Le(b, "0") == 0 {
		return b
	}
	n := toDec(a).Div(toDec(b)).Mul(toDec("100"))
	return n.RoundDown(3).String()
}

// ( (a-b)/b )*100
func Rose(a, b string) string {
	if Le(b, "0") == 0 {
		return b
	}
	n := toDec(a).Sub(toDec(b)).Div(toDec(b)).Mul(toDec("100"))
	return n.String()
}

// ( (a-b)/b )*100 保留 3 位小数
func RoseCent(a, b string) string {
	if Le(b, "0") == 0 {
		return b
	}
	n := toDec(a).Sub(toDec(b)).Div(toDec(b)).Mul(toDec("100"))
	return n.RoundDown(3).String()
}

// 比大小 a - b =  -1  0  1
func Le(a, b string) int {
	r := Sub(a, b)
	find := strings.Contains(r, "-")
	if r == "0" {
		return 0
	}
	if find {
		return -1
	} else {
		return 1
	}
}

// |a| 绝对值
func Abs(a string) string {
	n := toDec(a)
	h := n.Abs()
	return h.String()
}
