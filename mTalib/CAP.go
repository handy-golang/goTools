package mTalib

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
)

func CAP(opt CListOpt) string {
	KDList := opt.CList
	n := opt.Period
	cLen := len(KDList)

	if cLen < n {
		return "0"
	}

	last := KDList[cLen-1]
	start := KDList[cLen-n]

	// a-b 的涨幅 * 100 保留两位小数 然后 除以 n
	//

	diffC := mCount.Rose(last, start)
	diffRc := mCount.Mul(diffC, "100")
	diffCent := mCount.Cent(diffRc, 2)

	p := mCount.Mul(diffCent, fmt.Sprint(n-1))

	return p
}
