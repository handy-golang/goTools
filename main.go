package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mVerify"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	for i := 0; i < 10; i++ {
		vCode := mVerify.NewCode()
		fmt.Println(vCode)
	}
	fmt.Println(" =========   END   ========= ")
}
