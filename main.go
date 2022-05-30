package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mTikker"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	mTikker.NewTikker(mTikker.TikkerOpt{
		ShellContent: `
		mEcho "我爱你"
		`,
	}).Run()

	fmt.Println(" =========   END   ========= ")
}
