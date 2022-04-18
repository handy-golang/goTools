package mFeiShu

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mFetch"
	jsoniter "github.com/json-iterator/go"
)

type Opt struct {
	AppID     string
	AppSecret string
	AppType   string // 企业内 & 应用商店  company &  store
	Event     func(string, string)
}

type NewFeiShu struct {
	AppID       string
	AppSecret   string
	AppType     string
	AccessToken string
}

func New(opt Opt) *NewFeiShu {
	// 检查参数
	if len(opt.AppID) < 5 {
		errStr := fmt.Errorf("缺少 AppID 参数 文档: %+v", "https://open.feishu.cn/document/home/index")
		panic(errStr)
	}
	if len(opt.AppSecret) < 5 {
		errStr := fmt.Errorf("缺少 AppSecret 参数 文档: %+v", "https://open.feishu.cn/document/home/index")
		panic(errStr)
	}

	var o NewFeiShu
	o.AppID = opt.AppID
	o.AppSecret = opt.AppSecret

	if len(opt.AppType) < 5 {
		o.AppType = "company"
	} else {
		o.AppType = opt.AppType
	}

	// 获取 AccessToken
	mCycle.New(mCycle.CycleParam{
		Func: func() {
			o.GetAccessToken()
		},
		SleepTime: time.Hour,
	}).Start()

	return &o
}

type TenantAccessToken struct {
	Code              int    `json:"code"`
	Expire            int    `json:"expire"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

func (o *NewFeiShu) GetAccessToken() *NewFeiShu {
	data := map[string]any{
		"app_id":     o.AppID,
		"app_secret": o.AppSecret,
	}

	Path := "/open-apis/auth/v3/tenant_access_token/internal"

	if o.AppType == "store" {
		Path = "/open-apis/auth/v3/tenant_access_token"
	}

	res := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: "https://open.feishu.cn",
		Path:   Path,
		Data:   data,
	}).Post()

	if len(res) < 5 {
		return o
	}

	var result TenantAccessToken
	jsoniter.Unmarshal(res, &result)

	if result.Code == 0 {
		o.AccessToken = result.TenantAccessToken
	}

	return o
}
