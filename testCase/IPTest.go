package testCase

import (
	"github.com/EasyGolang/goTools/mVerify"
)

func IPTest() {
	ip := "36.44.232.38"
	mVerify.GetIPaddress(ip)
}
