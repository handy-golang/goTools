package testCase

import (
	"log"

	"github.com/EasyGolang/goTools/mVerify"
)

func IPTest() {
	ip := "36.44.232.38"
	resData, err := mVerify.GetIPaddress(ip)

	log.Println(resData, err)
}
