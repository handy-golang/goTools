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
		mEcho "我爱你22"
		`,
	}).RunToPm2()

	fmt.Println(" =========   END   ========= ")
}
