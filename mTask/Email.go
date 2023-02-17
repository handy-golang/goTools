package mTask

import (
	"github.com/EasyGolang/goTools/mJson"
	jsoniter "github.com/json-iterator/go"
)

// 验证码邮件
type CodeEmail struct {
	To         []string `bson:"To"`
	VerifyCode string   `bson:"VerifyCode"` // 验证码
	Action     string   `bson:"Action"`     // 行为
}

// 注册成功邮件
type RegisterParam struct {
	Password string `bson:"Password"` // 密码
}

// 邮件通用模板
type SendEmail struct {
	From     string     `bson:"From"`
	To       []string   `bson:"To"`
	Subject  string     `bson:"Subject"`
	TmplName string     `bson:"TmplName"` // 邮件模板的名字
	SendData EmailParam `bson:"SendData"` // 邮件模板需要的数据
}

// 邮件通用模板
type EmailParam struct {
	Title        string `bson:"Title"`   // 标题
	Message      string `bson:"Message"` // 消息
	Content      string `bson:"Content"` // 内容
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
