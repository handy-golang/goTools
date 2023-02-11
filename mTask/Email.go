package mTask

import (
	"github.com/EasyGolang/goTools/mJson"
	jsoniter "github.com/json-iterator/go"
)

// 系统邮件
type SysEmailParam struct {
	Title        string `bson:"Title"`        // 标题
	Message      string `bson:"Message"`      // 消息
	Content      string `bson:"Content"`      // 内容
	SysTime      string `bson:"SysTime"`      // 系统时间
	Source       string `bson:"Source"`       // 来源
	SecurityCode string `bson:"SecurityCode"` // 安全码
}
type SysEmail struct {
	From     string        `bson:"From"`     // 从哪里来
	To       []string      `bson:"To"`       // 发给谁
	Subject  string        `bson:"Subject"`  // 标题是啥
	SendData SysEmailParam `bson:"SendData"` // 要发送邮件的数据内容
}

// 验证码邮件
type CodeEmailParam struct {
	VerifyCode   string `bson:"VerifyCode"`
	Action       string `bson:"Action"`
	Minute       string `bson:"Minute"`       // 分钟数
	SysTime      string `bson:"SysTime"`      // 系统时间
	Source       string `bson:"Source"`       // 来源
	SecurityCode string `bson:"SecurityCode"` // 安全码
}
type CodeEmail struct {
	From     string         `bson:"From"`     // 从哪里来
	To       []string       `bson:"To"`       // 发给谁
	Subject  string         `bson:"Subject"`  // 标题是啥
	SendData CodeEmailParam `bson:"SendData"` // 要发送邮件的数据内容
}

// 注册成功邮件
type RegisterParam struct {
	Password     string `bson:"Password"`
	SysTime      string `bson:"SysTime"`      // 系统时间
	Source       string `bson:"Source"`       // 来源
	SecurityCode string `bson:"SecurityCode"` // 安全码
}
type RegisterEmail struct {
	From     string        `bson:"From"`     // 从哪里来
	To       []string      `bson:"To"`       // 发给谁
	Subject  string        `bson:"Subject"`  // 标题是啥
	SendData RegisterParam `bson:"SendData"` // 要发送邮件的数据内容
}

func ToMapData(val SysEmail) (resData map[string]any) {
	// 转换结果
	jsonStr := mJson.ToJson(val)
	jsoniter.Unmarshal(jsonStr, &resData)
	return
}
