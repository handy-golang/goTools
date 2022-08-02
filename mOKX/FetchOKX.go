package mOKX

import (
	"io/ioutil"
	"strings"

	"github.com/EasyGolang/goTools/mFetch"
)

type OptFetchOKX struct {
	Path          string
	Data          map[string]any
	LocalJsonPath string // 本地的数据源
	Method        string
	IsLocalJson   bool
	Event         func(string, any)
}

func FetchOKX(opt OptFetchOKX) (resData []byte, resErr error) {
	// 本地模式
	if opt.IsLocalJson {
		return ioutil.ReadFile(opt.LocalJsonPath)
	}

	if len(opt.Method) < 1 {
		opt.Method = "GET"
	}

	// 处理 Header 和 加密信息
	Method := strings.ToUpper(opt.Method)

	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "https://www.okx.com",
		Path:   opt.Path,
		Data:   opt.Data,
		Event:  opt.Event,
	})

	if Method == "GET" {
		return fetch.Get()
	} else {
		return fetch.Post()
	}
}
