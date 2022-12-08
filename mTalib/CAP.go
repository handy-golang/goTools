package mTalib

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
)

func CAP(opt CListOpt) string {
	KDList := opt.CList
	n := opt.Cycle

	c_len := len(KDList)
	last := KDList[c_len-1]
	if c_len < n {
		c_len = n
	}
	start := KDList[c_len-n]

	// a-b 的涨幅 * 100 保留两位小数 然后 除以 n
	//

	diffC := mCount.Rose(last, start)
	diffRc := mCount.Mul(diffC, "100")
	diffCent := mCount.Cent(diffRc, 2)

	p := mCount.Mul(diffCent, fmt.Sprint(n-1))

	return p
}
