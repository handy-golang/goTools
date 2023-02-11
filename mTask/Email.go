package mTask

import (
	"github.com/EasyGolang/goTools/mJson"
	jsoniter "github.com/json-iterator/go"
)

// 邮件任务需要的基本
type SysEmail struct {
	From     string   `bson:"From"`    // 从哪里来
	To       []string `bson:"To"`      // 发给谁
	Subject  string   `bson:"Subject"` // 标题是啥
	SendData struct {
		Title        string `bson:"Title"`        // 标题
		Message      string `bson:"Message"`      // 消息
		Content      string `bson:"Content"`      // 内容
		SysTime      string `bson:"SysTime"`      // 系统时间
		Source       string `bson:"Source"`       // 来源
		SecurityCode string `bson:"SecurityCode"` // 安全码
	} `bson:"SendData"` // 要发送邮件的数据内容
}

type CodeEmail struct {
	From     string   `bson:"From"`    // 从哪里来
	To       []string `bson:"To"`      // 发给谁
	Subject  string   `bson:"Subject"` // 标题是啥
	SendData struct {
		VerifyCode   string `bson:"VerifyCode"`
		Action       string `bson:"Action"`
		Minute       string `bson:"Minute"`       // 分钟数
		SysTime      string `bson:"SysTime"`      // 系统时间
		Source       string `bson:"Source"`       // 来源
		SecurityCode string `bson:"SecurityCode"` // 安全码
	} `bson:"SendData"` // 要发送邮件的数据内容
}

type RegisterEmail struct {
	From     string   `bson:"From"`    // 从哪里来
	To       []string `bson:"To"`      // 发给谁
	Subject  string   `bson:"Subject"` // 标题是啥
	SendData struct {
		Password     string `bson:"Password"`
		SysTime      string `bson:"SysTime"`      // 系统时间
		Source       string `bson:"Source"`       // 来源
		SecurityCode string `bson:"SecurityCode"` // 安全码
	} `bson:"SendData"` // 要发送邮件的数据内容
}

func ToMapData(val SysEmail) (resData map[string]any) {
	// 转换结果
	jsonStr := mJson.ToJson(val)
	jsoniter.Unmarshal(jsonStr, &resData)

	return
}
