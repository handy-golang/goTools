package mVerify

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

type VerifyCode struct {
	Code string `bson:"Code"`
	Time string `bson:"Time"`
}

func NewCode() VerifyCode {
	code := mCount.GetRound(100000, 999999)
	var vCode VerifyCode
	vCode.Code = mStr.ToStr(code)
	vCode.Time = mTime.GetUnix()
	return vCode
}
