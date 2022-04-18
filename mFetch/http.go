package mFetch

import (
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/mUrl"
	"github.com/gocolly/colly"
	jsoniter "github.com/json-iterator/go"
)

type HttpParam struct {
	Origin string
	Path   string
	Data   map[string]any
	Method string
	Header map[string]string
}

type Http struct {
	Url    string
	Data   []byte
	Header map[string]string
}

func NewHttp(opt HttpParam) []byte {
	// 参数的错误处理
	if len(opt.Method) < 2 {
		errStr := fmt.Errorf("缺少 Method 参数")
		panic(errStr)
	}
	if len(opt.Path) < 2 {
		errStr := fmt.Errorf("缺少 Path 参数")
		panic(errStr)
	}
	if len(opt.Origin) < 2 {
		errStr := fmt.Errorf("缺少 Origin 参数")
		panic(errStr)
	}

	var HttpO Http
	HttpO.Url = opt.Origin + opt.Path
	HttpO.Header = opt.Header

	if strings.ToLower(opt.Method) == "get" {
		// 处理参数
		urlO := mUrl.InitUrl(HttpO.Url)
		for key, val := range opt.Data {
			v := fmt.Sprintf("%+v", val)
			urlO.AddParam(key, v)
		}
		HttpO.Url = urlO.String()
		return HttpO.Get()

	}

	if strings.ToLower(opt.Method) == "post" {

		data, _ := jsoniter.Marshal(opt.Data)
		HttpO.Data = data

		return HttpO.Post()
	}

	return []byte("")
}

// GET
func (o *Http) Get() []byte {
	Url := o.Url

	fmt.Println("get请求", Url)

	var body []byte
	var resError []byte
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json; charset=utf-8")
		// 添加header头
		for key, val := range o.Header {
			r.Headers.Set(key, val)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		body = r.Body
	})
	c.OnError(func(r *colly.Response, err error) {
		resError = r.Body
	})

	c.Visit(Url)

	if len(body) > 0 {
		return body
	} else {
		return resError
	}
}

func (o *Http) Post() []byte {
	url := o.Url
	data := o.Data

	var body []byte
	var resError []byte
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json; charset=utf-8")
		// 添加header头
		for key, val := range o.Header {
			r.Headers.Set(key, val)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		body = r.Body
	})
	c.OnError(func(r *colly.Response, err error) {
		resError = r.Body
	})

	c.PostRaw(url, data)

	if len(body) > 0 {
		return body
	} else {
		return resError
	}
}
