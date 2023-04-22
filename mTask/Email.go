package mTask

// ====== 系统邮件 ========
type SysEmailParam struct {
	Title          string `bson:"Title"`   // 标题
	Message        string `bson:"Message"` // 消息
	Content        string `bson:"Content"` // 内容
	SysTime        string `bson:"SysTime"`
	Source         string `bson:"Source"`
	EntrapmentCode string `bson:"EntrapmentCode"` // 防钓鱼码
}

type SysEmail struct {
	To       []string      `bson:"To"` // To 可以是 UserID
	From     string        `bson:"From"`
	Subject  string        `bson:"Subject"`
	SendData SysEmailParam `bson:"SendData"` // 邮件模板需要的数据
}

// ====== 验证码邮件 ========
type CodeEmailParam struct {
	VerifyCode     string `bson:"VerifyCode"`
	Action         string `bson:"Action"`
	SysTime        string `bson:"SysTime"`
	Source         string `bson:"Source"`
	EntrapmentCode string `bson:"EntrapmentCode"` // 防钓鱼码
}

type CodeEmail struct {
	To       string         `bson:"To"` // 验证码任务一次只能是一个
	From     string         `bson:"From"`
	Subject  string         `bson:"Subject"`
	SendData CodeEmailParam `bson:"SendData"` // 邮件模板需要的数据
}

// ====== 注册成功邮件 ========
type RegisterSucceedEmailParam struct {
	Password       string `bson:"SysTime"`
	SysTime        string `bson:"SysTime"`
	Source         string `bson:"Source"`
	EntrapmentCode string `bson:"EntrapmentCode"` // 防钓鱼码
}

type RegisterSucceedEmail struct {
	To       string                    `bson:"To"` // 注册成功一次也只能是一个
	From     string                    `bson:"From"`
	Subject  string                    `bson:"Subject"`
	SendData RegisterSucceedEmailParam `bson:"SendData"` // 邮件模板需要的数据
}
