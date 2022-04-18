package mCount

// 转为 int
func ToInt(a string) int {
	s := toDec(a).IntPart()
	return int(s)
}

// 转为 float
func ToFloat(a string, x int32) float64 {
	s := toDec(a)
	var f float64
	if x > 0 {
		f = s.RoundDown(x).InexactFloat64()
	} else {
		f = s.InexactFloat64()
	}
	return f
}
