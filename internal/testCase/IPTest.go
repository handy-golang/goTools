package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mVerify"
)

func IPTest() {
	/*
		ip := "36.44.232.38"


		resData, err := mVerify.GetIPaddress(ip)
		fmt.Println(resData, err)
		res := mVerify.GetIPS([]string{ip})
		mJson.Println(res)

	*/

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60"
	res := mVerify.DeviceToUA(userAgent)

	fmt.Println(res)
}
