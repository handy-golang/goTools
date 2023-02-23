package mVerify

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type IPAddressType struct {
	ISP      string
	Hostname string
	Country  string
	Region   string
	City     string
}

// 利用 https://whatismyipaddress.com/ip/36.44.232.38
// 可能需要墙外才能访问

func GetIPS(ips []string) {
	// for _, val := range ips {
	// }
}

func GetIPaddress(ip string) (resData IPAddressType, resErr error) {
	if !IsIP(ip) {
		resErr = fmt.Errorf("ip地址不正确")
		return
	}

	var collyData []byte
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		// 添加header头
	})
	c.OnResponse(func(r *colly.Response) {
		collyData = r.Body
		fmt.Println(2222)
	})
	c.OnError(func(r *colly.Response, errStr error) {
		collyData = r.Body
		resErr = errStr

		fmt.Println(11111)
	})

	c.Visit("https://whatismyipaddress.com/ip/36.44.232.38")

	if resErr != nil {
		return
	}

	log.Println(string(collyData))

	return
}
