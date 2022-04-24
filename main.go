package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	for i := 0; i < 10; i++ {
		uuid := mEncrypt.GetUUID()
		fmt.Println(uuid)
	}
	fmt.Println(" =========   END   ========= ")
}
