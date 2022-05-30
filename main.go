package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mLog"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	for i := 0; i < 100; i++ {
		mLog.NewLog(mLog.NewLogParam{
			Name: "test",
		})
	}

	fmt.Println(" =========   END   ========= ")
}
