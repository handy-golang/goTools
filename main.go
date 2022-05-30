package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mTikker"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	mTikker.NewTikker(mTikker.TikkerOpt{
		ShellContent: "",
	}).InstPm2()

	fmt.Println(" =========   END   ========= ")
}
