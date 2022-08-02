package mOKX

import (
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/mTime"
)

/*
	wss := wssApi.New(wssApi.FetchOpt{
		Type: 1,
		Event: func(s string, a any) {
			global.WssLog.Println("Event", s, mStr.ToStr(a))
		},
	})

	go wss.Read(func(msg []byte) {
		global.WssLog.Println("读数据", mStr.ToStr(msg))
	})

*/

type LoginArgsType struct {
	APIKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}
type LoginType struct {
	Op   string          `json:"op"`
	Args []LoginArgsType `json:"args"`
}

type OptWssOKX struct {
	FetchType int
	Event     func(string, any)
	OKXKey    TypeOkxKey
}

func WssOKX(opt OptWssOKX) (_this *mFetch.Wss) {
	WssOpt := mFetch.WssOpt{}
	WssOpt.Event = opt.Event
	if opt.FetchType == 0 {
		WssOpt.Url = "wss://ws.okx.com:8443/ws/v5/public"
	}
	if opt.FetchType == 1 {
		WssOpt.Url = "wss://ws.okx.com:8443/ws/v5/private"
	}
	_this = mFetch.NewWss(WssOpt)

	if opt.FetchType == 1 {
		Timestamp := mTime.EpochTime()
		SignStr := mStr.Join(
			Timestamp,
			"GET",
			"/users/self/verify",
		)
		Sign := mEncrypt.Sha256(SignStr, opt.OKXKey.SecretKey)
		LoginParam := LoginType{
			Op: "login",
			Args: []LoginArgsType{
				{
					APIKey:     opt.OKXKey.ApiKey,
					Passphrase: opt.OKXKey.Passphrase,
					Timestamp:  Timestamp,
					Sign:       Sign,
				},
			},
		}
		_this.Write(mJson.ToJson(LoginParam))
	}

	return
}