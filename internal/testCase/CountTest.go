package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mCount"
)

func CountTest() {
	for i := 0; i < 10; i++ {
		num := mCount.GetRound(-5, 1)
		fmt.Println(num)
	}
}
