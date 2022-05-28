package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	for i := 0; i < 20; i++ {
		TimeID := mEncrypt.TimeID()

		fmt.Println(TimeID)
	}

	fmt.Println(" =========   END   ========= ")
}
