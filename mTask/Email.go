package mTask

import (
	"github.com/EasyGolang/goTools/mJson"
	jsoniter "github.com/json-iterator/go"
)

type SendEmail struct {
	From     string   `bson:"From"`
	To       []string `bson:"To"`
	Subject  string   `bson:"Subject"`
	TmplName string   `bson:"TmplName"` // 邮件模板的名字
	SendData any      `bson:"SendData"` // 邮件模板需要的数据
}

// 系统邮件
type SysEmailParam struct {
	Title        string `bson:"Title"`   // 标题
	Message      string `bson:"Message"` // 消息
	Content      string `bson:"Content"` // 内容
	SysTime      string `bson:"SysTime"`
	Source       string `bson:"Source"`
	SecurityCode string `bson:"SecurityCode"`
}

// 验证码邮件
type CodeEmailParam struct {
	VerifyCode   string `bson:"VerifyCode"` // 验证码
	Action       string `bson:"Action"`     // 行为
	SysTime      string `bson:"SysTime"`
	Source       string `bson:"Source"`
	SecurityCode string `bson:"SecurityCode"`
}

// 注册成功邮件
type RegisterParam struct {
	Password     string `bson:"Password"` // 密码
	SysTime      string `bson:"SysTime"`
	Source       string `bson:"Source"`
	SecurityCode string `bson:"SecurityCode"`
}

// 结构转换
func ToMapData(val SendEmail) (resData map[string]any) {
	jsonStr := mJson.ToJson(val)
	jsoniter.Unmarshal(jsonStr, &resData)
	return
}
