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
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mStr.ToStr(resData))
}
