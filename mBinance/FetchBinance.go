package mBinance

import (
	"fmt"
	"os"

	"github.com/EasyGolang/goTools/mPath"
)

type TypeBinanceKey struct {
	Name      string `bson:"Name"`
	ApiKey    string `bson:"ApiKey"`
	SecretKey string `bson:"SecretKey"`
	IsTrade   bool   `bson:"IsTrade"`
	UserID    string `bson:"UserID"`
}

type OptFetchBinance struct {
	Path          string
	Data          map[string]any
	Method        string
	IsLocalJson   bool
	LocalJsonPath string // 本地的数据源
	Event         func(string, any)
	BinanceKey    TypeBinanceKey
}

var BinanceBaseUrl = "https://fapi.binance.com"

func FetchBinance(opt OptFetchBinance) (resData []byte, resErr error) {
	// 是否为本地模式
	if opt.IsLocalJson {
		isJsonPath := mPath.Exists(opt.LocalJsonPath)
		if isJsonPath {
			return os.ReadFile(opt.LocalJsonPath)
		} else {
			resErr = fmt.Errorf("LocalJsonPath")
			return
		}
	}

	if len(opt.Method) < 1 {
		opt.Method = "GET"
	}

	fullURL := BinanceBaseUrl + opt.Path

	fmt.Println("fullURL", fullURL)

	// req, err := http.NewRequest(opt.Method, fullURL)

	return
}
