package mOKX

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
	"github.com/EasyGolang/goTools/mUrl"
)

/*

testCase.OKXFetch()

*/
type TypeOkxKey struct {
	ApiKey     string
	SecretKey  string
	Passphrase string
}
type OptFetchOKX struct {
	Path          string
	Data          map[string]any
	Method        string
	IsLocalJson   bool
	LocalJsonPath string // 本地的数据源
	Event         func(string, any)
	OKXKey        TypeOkxKey
}

func FetchOKX(opt OptFetchOKX) (resData []byte, resErr error) {
	// 是否为本地模式
	if opt.IsLocalJson {
		isJsonPath := mPath.Exists(opt.LocalJsonPath)
		if isJsonPath {
			return ioutil.ReadFile(opt.LocalJsonPath)
		} else {
			resErr = fmt.Errorf("LocalJsonPath")
			return
		}
	}

	if len(opt.Method) < 1 {
		opt.Method = "GET"
	}

	// 处理 Header 和 加密信息
	Method := strings.ToUpper(opt.Method)
	Timestamp := mTime.IsoTime(true)
	ApiKey := opt.OKXKey.ApiKey
	SecretKey := opt.OKXKey.SecretKey
	Passphrase := opt.OKXKey.Passphrase
	Body := mJson.ToJson(opt.Data)

	SignStr := mStr.Join(
		Timestamp,
		Method,
		opt.Path,
		string(Body),
	)

	if Method == "GET" {
		Body = []byte("")
		urlO := mUrl.InitUrl(opt.Path)
		for key, val := range opt.Data {
			v := fmt.Sprintf("%+v", val)
			urlO.AddParam(key, v)
		}
		signPath := urlO.String()
		SignStr = mStr.Join(
			Timestamp,
			Method,
			signPath,
			string(Body),
		)
	}
	Sign := mEncrypt.Sha256(SignStr, SecretKey)
	fetch := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "https://www.okx.com",
		Path:   opt.Path,
		Data:   opt.Data,
		Event:  opt.Event,
		Header: map[string]string{
			"OK-ACCESS-KEY":        ApiKey,
			"OK-ACCESS-SIGN":       Sign,
			"OK-ACCESS-TIMESTAMP":  Timestamp,
			"OK-ACCESS-PASSPHRASE": Passphrase,
		},
	})

	if Method == "GET" {
		return fetch.Get()
	} else {
		return fetch.Post()
	}
}
