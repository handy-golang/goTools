package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/testCase"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	testCase.OKXFetch()

	fmt.Println(" =========   END   ========= ")
}
