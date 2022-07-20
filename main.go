package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mStr"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	wssConn := mFetch.NewWss(mFetch.WssOpt{
		Url: "ws://127.0.0.1:8999/api/wss",
		Event: func(s string, a any) {
			if s == "close" || s == "err" {
				fmt.Println("出错了", mStr.ToStr(a))
			}
		},
	})

	wssConn.Write([]byte("123"))

	wssConn.Read(func(msg []byte) {
		fmt.Println("read", string(msg))
	})

	fmt.Println(" =========   END   ========= ")
}
