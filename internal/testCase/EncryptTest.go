package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func EncryptTest() {
	Key := mEncrypt.MD5("123123123") // e39a2a711f1c1646a9dfcd4526dc43fe
	got := mEncrypt.Decrypt("a0a34e98213e590878cf2ef50a108469", Key)
	fmt.Println(string(got))
}
