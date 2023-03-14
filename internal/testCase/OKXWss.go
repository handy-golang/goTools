package testCase

import (
	"github.com/EasyGolang/goTools/internal/global"
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
			ApiKey:     "xxxx",
			SecretKey:  "xxxx",
			Passphrase: "xxx",
		},
	})

	wss.Read(func(msg []byte) {
		global.WssLog.Println("读数据", mStr.ToStr(msg))
	})
}
