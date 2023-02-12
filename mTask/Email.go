package mTask

import (
	"github.com/EasyGolang/goTools/mJson"
	jsoniter "github.com/json-iterator/go"
)

// 系统邮件
type SysEmailParam struct {
	Title        string `bson:"Title"`   // 标题
	Message      string `bson:"Message"` // 消息
	Content      string `bson:"Content"` // 内容
	SysTime      string `bson:"SysTime"`
	Source       string `bson:"Source"`
	SecurityCode string `bson:"SecurityCode"`
}
type SysEmail struct {
	From     string        `bson:"From"`
	To       []string      `bson:"To"`
	Subject  string        `bson:"Subject"`
	SendData SysEmailParam `bson:"SendData"` // 要发送邮件的数据内容
}

// 验证码邮件
type CodeEmailParam struct {
	VerifyCode   string `bson:"VerifyCode"` // 验证码
	Action       string `bson:"Action"`     // 行为
	Minute       string `bson:"Minute"`     // 分钟数
	SysTime      string `bson:"SysTime"`
	Source       string `bson:"Source"`
	SecurityCode string `bson:"SecurityCode"`
}
type CodeEmail struct {
	From     string         `bson:"From"`
	To       []string       `bson:"To"`
	Subject  string         `bson:"Subject"`
	SendData CodeEmailParam `bson:"SendData"` // 要发送邮件的数据内容
}

// 注册成功邮件
type RegisterParam struct {
	Password     string `bson:"Password"` // 密码
	SysTime      string `bson:"SysTime"`
	Source       string `bson:"Source"`
	SecurityCode string `bson:"SecurityCode"`
}
type RegisterEmail struct {
	From     string        `bson:"From"`
	To       []string      `bson:"To"`
	Subject  string        `bson:"Subject"`
	SendData RegisterParam `bson:"SendData"` // 要发送邮件的数据内容
}

func ToMapData(val SysEmail) (resData map[string]any) {
	// 转换结果
	jsonStr := mJson.ToJson(val)
	jsoniter.Unmarshal(jsonStr, &resData)
	return
}
