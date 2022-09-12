package mFile

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

type TinifyType struct {
	Input struct {
		Size int    `json:"size"`
		Type string `json:"type"`
	} `json:"input"`
	Output struct {
		Size   int     `json:"size"`
		Type   string  `json:"type"`
		Width  int     `json:"width"`
		Height int     `json:"height"`
		Ratio  float64 `json:"ratio"`
		URL    string  `json:"url"`
	} `json:"output"`
}

const CompressingUrl = "https://api.tinify.com/shrink"

// https://tinypng.com/developers
type TinyOpt struct {
	Src    string // 图片地址
	Email  string
	ApiKey string
}

func Tinypng(opt TinyOpt) (resData string, resErr error) {
	isPath := mPath.Exists(opt.Src)
	isFile := mPath.IsFile(opt.Src)

	if isPath && isFile {
	} else {
		resErr = fmt.Errorf("请输入有效的 Src")
		return
	}

	if len(opt.Email) < 6 || len(opt.ApiKey) < 10 {
		resErr = fmt.Errorf("请输入有效的 opt.Email 或 opt.ApiKey")
		return
	}

	extName := path.Ext(opt.Src) // 后缀名

	if extName == ".png" || extName == ".jpg" || extName == ".jpeg" || extName == ".webp" {
	} else {
		resErr = fmt.Errorf("文件格式不正确")
		return
	}

	// 创建Request
	req, err := http.NewRequest(http.MethodPost, CompressingUrl, nil)
	if err != nil {
		resErr = err
		return
	}
	req.SetBasicAuth(opt.Email, opt.ApiKey)

	data, err := os.ReadFile(opt.Src)
	if err != nil {
		resErr = err
		return
	}
	req.Body = io.NopCloser(bytes.NewReader(data))

	// 发起请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		resErr = err
		return
	}

	// 解析请求
	data, err = io.ReadAll(response.Body)
	if err != nil {
		resErr = err
		return
	}

	contStr := mStr.ToStr(data)
	find := strings.Contains(contStr, "output")
	if !find {
		resErr = fmt.Errorf(contStr)
		return
	}

	var result TinifyType
	jsoniter.Unmarshal(data, &result)

	if len(result.Output.URL) < 22 {
		resErr = fmt.Errorf("图片链接长度不足")
		return
	}
	resData = result.Output.URL
	return
}
