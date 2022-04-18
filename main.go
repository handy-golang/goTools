package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mFetch"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	Channel := "candle1m"
	InstID := "AVAX-USDT-mcl"

	data := map[string]any{
		"op": "subscribe",
		"args": []string{
			Channel, InstID, "mei",
		},
	}

	res := mFetch.NewHttp(mFetch.HttpParam{
		Origin: "http://mo7.cc:9000",
		Path:   "/api/markets/tickers",
		Data:   data,
		Method: "get",
		Header: map[string]string{
			"Content-Type":  "appli1arset=utf-8",
			"Content-Type1": "appl2et=utf-8",
			"Content-Type2": "applicati3et=utf-8",
			"Content-Type3": "application4t=utf-8",
		},
	})

	fmt.Println(string(res))

	fmt.Println(" =========   END   ========= ")
}
