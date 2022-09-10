package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mStr"
)

func StrFuzzy() {
	str := mStr.Fuzzy("15309140739")

	fmt.Println(str)
}
