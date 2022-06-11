package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	fmt.Println(mEncrypt.TimeID())

	fmt.Println(" =========   END   ========= ")
}
