package mFeiShu

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mFetch"
)

type Opt struct {
	AppID     string
	AppSecret string
	AppType   string // 企业内 & 应用商店  company &  store
	Event     func(string, string)
}

type NewFeiShu struct {
	AppID              string
	AppSecret          string
	AppType            string
	Getaccess_tokenUrl string
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
	o.AppType = opt.AppType

	if len(o.AppType) < 5 {
		o.AppType = "company"
	}

	// 获取 AccessToken
	mCycle.New(mCycle.CycleParam{
		Func: func() {
			o.GetAccessToken()
		},
		SleepTime: time.Second * 5,
	}).Start()

	return &o
}

func (o *NewFeiShu) GetAccessToken() *NewFeiShu {
	data := map[string]any{
		"app_id":     o.AppID,
		"app_secret": o.AppSecret,
	}

	Path := "open-apis/auth/v3/tenant_access_token/internal"

	res := mFetch.NewHttp(mFetch.HttpParam{
		Origin: "https://open.feishu.cn/open-apis",
		Path:   Path,
		Data:   data,
	}).Post()

	fmt.Println(string(res))

	return o
}
