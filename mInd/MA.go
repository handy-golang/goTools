package mInd

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

// 数据来源 (H+L+C)  / 3
func MA(KDList []mOKX.TypeKd, n int) string {
	c_len := len(KDList) // K线总长
	c_n := n             // 长度
	if c_len < n {
		c_n = c_len
	}

	c_list := KDList[c_len-c_n:]
	ma_add := "0"
	for _, KD := range c_list {
		C := KD.CBas
		ma_add = mCount.Add(ma_add, C)
	}

	maRe := mCount.Div(ma_add, mStr.ToStr(c_n))

	return maRe
}
