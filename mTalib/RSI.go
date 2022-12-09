package mTalib

import (
	"github.com/EasyGolang/goTools/mTalib/talib"
)

func (_this *ClistObj) RSI() *ClistObj {
	if _this.CLen < _this.Period+1 {
		return _this
	}
	pArr := talib.Rsi(_this.FList, _this.Period)
	_this.Result = pArr[_this.CLen-1]
	return _this
}
