package mVerify

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
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

	/* var collyData []byte
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50")
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
	*/
	collyUseTemplate()
	return
}

func collyUseTemplate() {
	// 创建采集器对象
	collector := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}), // 开启debug
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50"),
		colly.IgnoreRobotsTxt(), // 忽略目标机器中的`robots.txt`声明
	)
	// 发起请求之前调用
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("发起请求之前调用...")
	})
	// 请求期间发生错误,则调用
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("请求期间发生错误,则调用:", err)
	})
	// 收到响应后调用
	collector.OnResponse(func(response *colly.Response) {
		fmt.Println("收到响应后调用:", string(response.Body))
	})

	// url：请求具体的地址
	err := collector.Visit("https://whatismyipaddress.com")
	if err != nil {
		fmt.Println("具体错误:", err)
	}
}
