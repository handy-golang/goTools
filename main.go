package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mFeiShu"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	mFeiShu.New(mFeiShu.Opt{
		AppID:     "cli_a28394cb5478d00d",
		AppSecret: "MDMJs33KsiH9FAxr74MqSXG3lTL4ptPT",
	})

	fmt.Println(" =========   END   ========= ")
}
