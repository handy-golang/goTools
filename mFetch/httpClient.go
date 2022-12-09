package mFetch

import (
	"fmt"

	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mUrl"
	"github.com/gocolly/colly/v2"
	jsoniter "github.com/json-iterator/go"
)

/*

	data := map[string]any{
		"op": "subscribe",
		"args": []string{
			"123", "345", "mei",
		},
	}

	resData, err := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "http://localhost:9000",
		Path:   "/api/ping",
		Data:   data,
		Header: map[string]string{
			"Content-Type":  "appli1arset=utf-8",
			"Content-Type1": "appl2et=utf-8",
			"Content-Type2": "applicati3et=utf-8",
			"Content-Type3": "application4t=utf-8",
		},
	}).Get()
	if err != nil {
		fmt.Println("err", err)
	}
	jsonStr := mJson.JsonFormat(resData)
	fmt.Println("resData", jsonStr)

*/

type HttpOpt struct {
	Origin string
	Path   string
	Data   []byte
	Header map[string]string
	Event  func(string, any) // s1 = succeed , err
}

type Http struct {
	Url    string
	Data   []byte
	Header map[string]string
	Event  func(string, any)
}

func NewHttp(opt HttpOpt) (_this *Http) {
	_this = &Http{}
	// 检查参数
	errStr := []string{}
	switch {
	case len(opt.Origin) < 1:
		errStr = append(errStr, "Origin")
		fallthrough
	case len(opt.Path) < 1:
		errStr = append(errStr, "Path")
	}
	if len(errStr) > 0 {
		errStr := fmt.Errorf("缺少参数:%+v", errStr)
		panic(errStr)
	}

	_this.Url = mStr.Join(opt.Origin, opt.Path)
	_this.Data = opt.Data
	_this.Header = opt.Header
	_this.Event = opt.Event
	if _this.Event == nil {
		_this.Event = func(s1 string, s2 any) {}
	}

	return
}

// 处理 Get 参数
func (_this *Http) DisposeGetParam() *Http {
	urlO := mUrl.InitUrl(_this.Url)

	var dataUn map[string]any
	jsoniter.Unmarshal(_this.Data, &dataUn)

	for key, val := range dataUn {
		v := fmt.Sprintf("%+v", val)
		urlO.AddParam(key, v)
	}
	_this.Url = urlO.String()
	return _this
}

// GET
func (_this *Http) Get() (resData []byte, resErr error) {
	// 处理参数请求
	_this.DisposeGetParam()

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json; charset=utf-8")
		r.Headers.Set("User-Agent", "goTools - github.com/EasyGolang/goTools")
		// 添加header头
		for key, val := range _this.Header {
			r.Headers.Set(key, val)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		resData = r.Body
		_this.Event("succeed", resData)
	})
	c.OnError(func(r *colly.Response, errStr error) {
		resData = r.Body
		resErr = errStr
		_this.Event("err", errStr)
	})

	c.Visit(_this.Url)

	return
}

// Post

func (_this *Http) Post() (resData []byte, resErr error) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/json; charset=utf-8")
		r.Headers.Set("User-Agent", "goTools - github.com/EasyGolang/goTools")
		// 添加header头
		for key, val := range _this.Header {
			r.Headers.Set(key, val)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		resData = r.Body
		_this.Event("succeed", resData)
	})
	c.OnError(func(r *colly.Response, err error) {
		resData = r.Body
		resErr = err
		_this.Event("err", err)
	})

	c.PostRaw(_this.Url, _this.Data)

	return
}
