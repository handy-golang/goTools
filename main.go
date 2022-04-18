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
	})

	fmt.Println(string(res))

	fmt.Println(" =========   END   ========= ")
}
