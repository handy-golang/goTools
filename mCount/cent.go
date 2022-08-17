package mCount

import (
	"strings"
)

// 判断小数位数
func GetDecimal(str string) int {
	str_arr := strings.Split(str, ".")

	if len(str_arr) > 1 {
		dec := str_arr[1]
		return len(dec)
	} else {
		return 0
	}
}

// 保留 x 位小数，其余都舍弃
func Cent(a string, x int32) string {
	n := toDec(a)
	return n.RoundDown(x).String()
}

// 四舍五入 保留 x 位小数
func CentRound(a string, x int32) string {
	n := toDec(a)
	return n.Round(x).String()
}

// 按照 source 的小数位数来保留小数
func PriceCent(target, source string) string {
	decimalLen := GetDecimal(source)
	v := CentRound(target, int32(decimalLen))
	return v
}
