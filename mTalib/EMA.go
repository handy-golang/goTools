package mTalib

import (
	"github.com/EasyGolang/goTools/mTalib/talib"
)

func (_this *ClistObj) EMA() *ClistObj {
	if _this.CLen < _this.Period+1 {
		return _this
	}

	pArr := talib.Ema(_this.FList, _this.Period)
	_this.Result = pArr[_this.CLen-1]
	return _this
}

/**

// 这个库废弃了, 用table 重写

func EMA(opt CListOpt) string {
	KDList := opt.CList
	n := opt.Cycle

	c_len := len(KDList) // K线总长
	c_n := n             // 长度
	if c_len < n {
		c_n = c_len
	}

	y_list := KDList[0:c_n] // 将最开始的N个KD 作为初始参数
	y := MA(CListOpt{       // 初始值计算
		CList:     y_list,
		Cycle:     c_n,
		Precision: opt.Precision,
	})

	ema_list := KDList[c_n:]

	if len(opt.Precision) < 1 {
		opt.Precision = KDList[0]
	}
	dotNum := mCount.GetDecimal(opt.Precision) // 计算小数点位数

	for _, KD := range ema_list {
		C := KD

		tody := C                // 今日的价格
		q := "2"                 // 2* tody
		w := fmt.Sprint(c_n + 1) // n +1
		e := fmt.Sprint(c_n - 1) // n -1
		// 昨日 EMA 是 y
		t1 := mCount.Mul(q, tody) // 2* 今日收盘价
		u1 := mCount.Div(t1, w)   //  !!  2* 今日收盘价 /( 12+1 )
		t2 := mCount.Mul(e, y)    // (12-1) * 昨日 ema(12)
		u2 := mCount.Div(t2, w)   //!!  (12-1) * 昨日 ema(12)  / （12+1）
		y = mCount.Add(u1, u2)

		y = mCount.CentRound(y, dotNum)
	}

	return y
}


*/
