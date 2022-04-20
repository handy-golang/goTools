package main

import (
	"fmt"
	"os"

	"github.com/EasyGolang/goTools/mStr"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	// a := []rune("mo7欢迎你")
	// a := []byte("mo7欢迎你")
	// a := 10.97
	a := os.PathSeparator

	str := mStr.ToStr(a)

	fmt.Println(str)

	fmt.Println(" =========   END   ========= ")
}
