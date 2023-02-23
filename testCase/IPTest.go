package testCase

import (
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mVerify"
)

func IPTest() {
	ip := "36.44.232.38"
	// resData, err := mVerify.GetIPaddress(ip)
	// fmt.Println(resData, err)

	res := mVerify.GetIPS([]string{ip})
	mJson.Println(res)
}
