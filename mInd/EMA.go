package mInd

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
)

// 数据来源 (H+L+C) / 3
func EMA(KDList []mOKX.TypeKd, n int) string {
	c_len := len(KDList) // K线总长
	c_n := n             // 长度
	if c_len < n {
		c_n = c_len
	}
	y_list := KDList[0:c_n] // 将最开始的N个KD 作为初始参数
	y := MA(y_list, c_n)    // 初始值

	ema_list := KDList[n:]
	var precision int32

	for _, KD := range ema_list {
		C := KD.CBas

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

		// 精度
		precision = mCount.GetDecimal(KD.TickSz)
		y = mCount.CentRound(y, precision)
	}

	return y
}
