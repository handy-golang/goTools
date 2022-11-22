package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStr"
)

// OKX 加密请求
func OKXFetch() {
	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/account/balance",
		Data: map[string]any{
			"ccy": "USDT",
		},
		Method: "get",
		Event: func(s string, a any) {
			global.Log.Println("Event", s, a)
		},
		OKXKey: mOKX.TypeOkxKey{
			ApiKey:     "ca6e399d-cc78-41b0-90fe-40ab99fb8040",
			SecretKey:  "EF88B67D7E618FF6B854359BE88A6445",
			Passphrase: "@Mcl931750",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mStr.ToStr(resData))
}

func BinanceFetch() {
	resData, err := mOKX.FetchOKX(mOKX.OptFetchOKX{
		Path: "/api/v5/account/balance",
		Data: map[string]any{
			"ccy": "USDT",
		},
		Method: "get",
		Event: func(s string, a any) {
			global.Log.Println("Event", s, a)
		},
		OKXKey: mOKX.TypeOkxKey{
			ApiKey:     "ca6e399d-cc78-41b0-90fe-40ab99fb8040",
			SecretKey:  "EF88B67D7E618FF6B854359BE88A6445",
			Passphrase: "@Mcl931750",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mStr.ToStr(resData))
}
