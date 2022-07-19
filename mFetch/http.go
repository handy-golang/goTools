package mFetch

import (
	"fmt"

	"github.com/EasyGolang/goTools/mUrl"
	"github.com/gocolly/colly"
	jsoniter "github.com/json-iterator/go"
)

/*

	data := map[string]any{
		"op": "subscribe",
		"args": []string{
			Channel, InstID, "mei",
		},
	}

	res := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "http://mo7.cc:9000",
		Path:   "/api/markets/tickers",
		Data:   data,
		Header: map[string]string{
			"Content-Type":  "appli1arset=utf-8",
			"Content-Type1": "appl2et=utf-8",
			"Content-Type2": "applicati3et=utf-8",
			"Content-Type3": "application4t=utf-8",
		},
	}).Get()

*/

type Http struct {
	Url     string
	OptData map[string]any
	Data    []byte
	Header  map[string]string
	Event   func(string, string)
}

type HttpOpt struct {
	Origin string
	Path   string
	Data   map[string]any
	Header map[string]string
	Event  func(string, string)
}

func NewHttp(opt HttpOpt) *Http {
	var HttpO Http

	// 参数的错误处理
	if len(opt.Path) < 2 {
		errStr := fmt.Errorf("缺少 Path 参数")
		panic(errStr)
	}

	HttpO.Url = opt.Origin + opt.Path
	HttpO.Header = opt.Header

	// 函数空指针的处理
	if opt.Event != nil {
		HttpO.Event = opt.Event
	} else {
		HttpO.Event = func(s1, s2 string) {}
	}

	HttpO.OptData = opt.Data

	return &HttpO
}

func proessHttpParam(lType string, o *Http) *Http {
	if lType == "get" {
		urlO := mUrl.InitUrl(o.Url)
		for key, val := range o.OptData {
			v := fmt.Sprintf("%+v", val)
			urlO.AddParam(key, v)
		}
		o.Url = urlO.String()
	}

	if lType == "post" {
		data, _ := jsoniter.Marshal(o.OptData)
		o.Data = data
	}

	return o
}

// GET
func (o *Http) Get() []byte {
	// 处理参数请求
	proessHttpParam("get", o)

	Url := o.Url

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
		o.Event("succeed", string(r.Body))
	})
	c.OnError(func(r *colly.Response, _ error) {
		resError = r.Body
		o.Event("err", string(r.Body))
	})

	c.Visit(Url)

	if len(body) > 0 {
		return body
	} else {
		return resError
	}
}

func (o *Http) Post() []byte {
	proessHttpParam("post", o)

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
		o.Event("succeed", string(r.Body))
	})
	c.OnError(func(r *colly.Response, _ error) {
		resError = r.Body
		o.Event("err", string(r.Body))
	})

	c.PostRaw(url, data)

	if len(body) > 0 {
		return body
	} else {
		return resError
	}
}
