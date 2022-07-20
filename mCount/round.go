package mCount

import (
	"crypto/rand"
	"math"
	"math/big"
)

/*
	随机数整数
*/
func GetRound(min, max int64) int64 {
	if min > max {
		return 0
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
