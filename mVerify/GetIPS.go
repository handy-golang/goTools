package mVerify

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type IPAddressType struct {
	ISP      string
	Hostname string
	Country  string
	Region   string
	City     string
}

func GetIPS(ips []string) {
	// for _, val := range ips {
	// }
}

//go:embed WhatIsMyIpHeader.yaml
var WhatIsMyIpHeader string

/*
https://www.ipshudi.com/36.44.232.38.htm
*/
func GetIPaddress(ip string) (resData IPAddressType, resErr error) {
	if !IsIP(ip) {
		resErr = fmt.Errorf("ip地址不正确")
		return
	}

	HeaderMap := FileToHeader(WhatIsMyIpHeader)

	var collyData []byte
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		for key, val := range HeaderMap {
			r.Headers.Set(key, val)
		}
	})
	c.OnResponse(func(r *colly.Response) {
		collyData = r.Body
	})
	c.OnError(func(r *colly.Response, errStr error) {
		collyData = r.Body
		resErr = errStr
	})

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	if err := e.Request.Visit(e.Attr("href")); err != nil {
	// 		log.Printf("visit err: %v", err)
	// 	}
	// })

	c.Visit("https://www.ipshudi.com/36.44.232.38.htm")

	if resErr != nil {
		return
	}
	log.Println("请求的Body", resErr, string(collyData))

	return
}

func FileToHeader(cont string) map[string]string {
	strArr := strings.Split(cont, "\n")
	HeaderMap := make(map[string]string)
	for _, item := range strArr {
		kvArr := strings.Split(item, ": ")
		if len(kvArr) == 2 {
			k := kvArr[0]
			v := kvArr[1]
			HeaderMap[k] = v
		}
	}
	return HeaderMap
}
