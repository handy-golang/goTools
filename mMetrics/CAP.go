package mMetrics

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
)

func CAP(opt EmaOpt) string {
	KDList := opt.CList
	n := opt.Cycle

	c_len := len(KDList)
	last := KDList[c_len-1]
	if c_len < n {
		c_len = n
	}
	start := KDList[c_len-n]

	diffPrice := mCount.Sub(last, start)

	p := mCount.Mul(diffPrice, fmt.Sprint((n)))

	return p
}
