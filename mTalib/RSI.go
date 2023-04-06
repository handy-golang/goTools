package mTalib

import "github.com/EasyGolang/goTools/mTalib/talib"

func (_this *ClistObj) RSI() *ClistObj {
	if _this.CLen < _this.Period+1 {
		return _this
	}

	_this.DotNum = 4 // 固定保留2位小数
	pArr := talib.Rsi(_this.FList, _this.Period)
	_this.Result = pArr[_this.CLen-1]

	return _this
}
