package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mVerify"
)

func StrFuzzy() {
	// str := mStr.Fuzzy("15309140739")

	// str := mStr.GetKeyFuzzyStart("123456789", 3)
	str := mStr.GetKeyFuzzy("123456789", 3, 5)

	fmt.Println(str)
}

func IPtest() {
	str := "255.255.240.9"

	r := mVerify.IsIP(str)
	// r := mVerify.IsPort(str)

	fmt.Println(r)
}
