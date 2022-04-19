package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	ll := []string{
		"1", "2", "3", "4",
	}
	r := mCount.Sum(ll)
	fmt.Println(r)

	fmt.Println(" =========   END   ========= ")
}
