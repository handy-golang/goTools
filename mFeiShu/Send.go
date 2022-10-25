package mFeiShu

import (
	"fmt"

	"github.com/EasyGolang/goTools/mFetch"
	"github.com/EasyGolang/goTools/mJson"
)

type MsgOpt struct {
	ReceiveType string
	ReceiveId   string
	MsgType     string
	Content     string
}

func (o *NewFeiShu) SendMessage(opt MsgOpt) []byte {
	receive_id_type := opt.ReceiveType
	receive_id := opt.ReceiveId
	content := opt.Content
	msg_type := opt.MsgType

	if len(opt.ReceiveType) < 2 {
		fmt.Printf("缺少参数 ReceiveType , receive_id_type 阅读文档 : %+v \n", "https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create")
		return nil
	}
	if len(opt.ReceiveId) < 2 {
		fmt.Printf("缺少参数 ReceiveId , receive_id 阅读文档 : %+v \n", "https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/create")
		return nil
	}
	if len(opt.MsgType) < 2 {
		msg_type = "interactive"
	}

	if len(opt.Content) < 1 {
		content = Msg_1_card("Hello, goTools")
	}

	if msg_type == "interactive" {
		content = Msg_1_card(content)
	}

	Path := "/open-apis/im/v1/messages?receive_id_type=" + receive_id_type

	data := map[string]any{
		"receive_id": receive_id,
		"content":    content,
		"msg_type":   msg_type,
	}
	res, _ := mFetch.NewHttp(mFetch.HttpOpt{
		Origin: o.Origin,
		Path:   Path,
		Data:   mJson.ToJson(data),
		Header: map[string]string{
			"Content-Type":  "application/json; charset=utf-8",
			"Authorization": "Bearer " + o.AccessToken,
		},
	}).Post()

	o.Event("SendMessage", string(res))

	return res
}
