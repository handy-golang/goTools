package mUrl

import (
	"net/url"
)

type UrlPar struct {
	Url *url.URL
}

/*
	解构 url
	获取 url 参数
	使用方法：

	tools.ParseUrl("https://mo7.cc?abc=123")
*/
func InitUrl(urlStr string) *UrlPar {
	var urlInfo UrlPar
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	urlInfo.Url = u

	return &urlInfo
}

// 格式化成为字符串
func (o *UrlPar) String() string {
	str := o.Url.String()
	return str
}

// 获取 Url 的参数
func (o *UrlPar) GetParam(key string) string {
	sum, err := url.ParseQuery(o.Url.RawQuery)
	if err != nil {
		panic(err)
	}

	value := sum.Get(key)

	return value
}

func (o *UrlPar) ParseQuery() url.Values {
	sum, err := url.ParseQuery(o.Url.RawQuery)
	if err != nil {
		panic(err)
	}
	return sum
}

func (o *UrlPar) AddParam(key string, val string) *UrlPar {
	sum, err := url.ParseQuery(o.Url.RawQuery)
	if err != nil {
		panic(err)
	}

	if len(key) > 0 && len(val) > 0 {
		sum.Add(key, val)
		o.Url.RawQuery = sum.Encode()

		return o
	}

	return o
}
