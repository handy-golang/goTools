package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func EncryptTest() {
	got := mEncrypt.AseDecrypt("c5408fbfd8db878e888eed29330b4698", "abc123")
	fmt.Println(string(got))
}
