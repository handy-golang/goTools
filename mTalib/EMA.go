package mTalib

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mTalib/talib"
)

type CListOpt struct {
	CList  []string // 数据
	Period int      // 周期
}

func EMA(opt CListOpt) string {
	n := opt.Period
	cLen := len(opt.CList)

	if cLen < n {
		return "0"
	}

	dotNum := mCount.GetDecimal(opt.CList[0]) // 计算小数点位数
	var floatList []float64
	for _, val := range opt.CList {
		floatList = append(floatList, mCount.ToFloat(val, -1)) // 将数值完整的转化
		valDot := mCount.GetDecimal(opt.CList[0])              // 计算当前的小数点位数
		if valDot-dotNum > 0 {                                 // 如果当前小数点位数大于现存小数点位数，则替换
			dotNum = valDot
		}
	}
	// 计算 Ema 指标
	pArr := talib.Ema(floatList, n)
	emaFloat := pArr[cLen-1]

	// 保留精确度，并转为字符串
	emaStr := fmt.Sprintf("%f", emaFloat)
	emaStr = mCount.CentRound(emaStr, dotNum)

	return emaStr
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
