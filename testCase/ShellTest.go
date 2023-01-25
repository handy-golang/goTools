package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mShell"
)

func ShellTest() {
	res, err := mShell.Run(`
	echo "123"	
	`)

	fmt.Println(string(res), err)
}
