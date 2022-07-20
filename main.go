package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	data := map[string]any{
		"op": "subscribe",
		"args": []string{
			"123", "345", "mei",
		},
	}

	resData, err := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "http://localhost:9000",
		Path:   "/api/ping",
		Data:   data,
		Header: map[string]string{
			"Content-Type":  "appli1arset=utf-8",
			"Content-Type1": "appl2et=utf-8",
			"Content-Type2": "applicati3et=utf-8",
			"Content-Type3": "application4t=utf-8",
		},
	}).Get()
	if err != nil {
		fmt.Println("err", err)
	}
	jsonStr := mJson.JsonFormat(resData)
	fmt.Println("resData", jsonStr)

	fmt.Println(" =========   END   ========= ")
}
