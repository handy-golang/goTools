package mCount

import (
	"strings"
)

// 判断小数位数
func GetDecimal(str string) int32 {
	str_arr := strings.Split(str, ".")

	if len(str_arr) > 1 {
		dec := str_arr[1]
		return int32(len(dec))
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
