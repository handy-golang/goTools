package mVerify

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
)

func NewCode() string {
	code := mCount.GetRound(100000, 999999)
	return mStr.ToStr(code)
}
