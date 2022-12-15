package mTalib

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
)

func (_this *ClistObj) CAP() *ClistObj {
	if _this.CLen < _this.Period+1 {
		return _this
	}

	last := _this.CList[_this.CLen-1]
	start := _this.CList[_this.CLen-_this.Period]

	// a-b 的涨幅 * 100 保留两位小数 然后 除以 n
	diffC := mCount.Rose(last, start)
	diffRc := mCount.Mul(diffC, "100")
	p := mCount.Div(diffRc, fmt.Sprint(_this.Period-1))

	floatVal := mCount.ToFloat(p, 2)
	_this.Result = floatVal

	return _this
}
