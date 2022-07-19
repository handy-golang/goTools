package mFeiShu

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mCycle"
	"github.com/EasyGolang/goTools/mFetch"
	jsoniter "github.com/json-iterator/go"
)

/*

	feishuApp := mFeiShu.New(mFeiShu.Opt{
		AppID:     "cli_a2xxxxxd00d",
		AppSecret: "MDMJxxxxxxlTL4ptPT",
	})

	str := mStr.Join(
		"交易方向: **", "开多", "** \n",
		"交易币种: **", "avax", "** \n",
	)

	feishuApp.SendMessage(mFeiShu.MsgOpt{
		ReceiveType: "user_id",
		ReceiveId:   "d8xxxxgc",
		Content:     str,
	})

*/

type NewFeiShu struct {
	Origin      string
	AppID       string
	AppSecret   string
	CardType    int
	AppType     string
	AccessToken string
	Event       func(string, string)
}

type Opt struct {
	AppID     string
	AppSecret string
	AppType   string // 企业内 & 应用商店  company &  store
	CardType  int    // 1 ,2 ,3, 4
	Event     func(string, string)
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
	o.Origin = "https://open.feishu.cn"

	if opt.CardType < 1 {
		o.CardType = 1
	} else {
		o.CardType = opt.CardType
	}

	// 函数空指针的处理
	if opt.Event != nil {
		o.Event = opt.Event
	} else {
		o.Event = func(s1, s2 string) {}
	}

	if len(opt.AppType) < 5 {
		o.AppType = "company"
	} else {
		o.AppType = opt.AppType
	}

	// 获取 AccessToken
	mCycle.New(mCycle.Opt{
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

	res, err := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: o.Origin,
		Path:   Path,
		Data:   data,
	}).Post()
	if err != nil {
		return o
	}

	o.Event("GetAccessToken", string(res))

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
