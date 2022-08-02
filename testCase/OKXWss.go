package testCase

import (
	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

func OKXWss() {
	wss := mOKX.WssOKX(mOKX.OptWssOKX{
		FetchType: 1,
		Event: func(s string, a any) {
			global.WssLog.Println("Event", s, mStr.ToStr(a))
		},
		OKXKey: mOKX.TypeOkxKey{
			ApiKey:     "ca6e399d-cc78-41b0-90fe-40ab99fb8040",
			SecretKey:  "EF88B67D7E618FF6B854359BE88A6445",
			Passphrase: "@Mcl931750",
		},
	})

	wss.Read(func(msg []byte) {
		global.WssLog.Println("读数据", mStr.ToStr(msg))
	})
}
