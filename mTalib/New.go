package mTalib

import "github.com/EasyGolang/goTools/mOKX"

type ClistOpt struct {
	CList  []string      // 数据
	KDList []mOKX.TypeKd // 数据
	Period int           // 周期
}

type ClistObj struct{}

func New() *ClistObj {
	obj := ClistObj{}

	return &obj
}
