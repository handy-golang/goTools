package mTask

// 邮件任务需要的基本
type SysEmail struct {
	From     string   `bson:"From"`    // 从哪里来
	To       []string `bson:"To"`      // 发给谁
	Subject  string   `bson:"Subject"` // 标题是啥
	SendData struct {
		Title        string // 标题
		Message      string // 消息
		Content      string // 内容
		SysTime      string // 系统时间
		Source       string // 来源
		SecurityCode string // 安全码
	} `bson:"SendData"` // 要发送邮件的数据内容
}

type CodeEmail struct {
	From     string   `bson:"From"`    // 从哪里来
	To       []string `bson:"To"`      // 发给谁
	Subject  string   `bson:"Subject"` // 标题是啥
	SendData struct {
		VerifyCode   string
		Action       string
		Minute       string // 分钟数
		SysTime      string // 系统时间
		Source       string // 来源
		SecurityCode string // 安全码
	} `bson:"SendData"` // 要发送邮件的数据内容
}

type RegisterEmail struct {
	From     string   `bson:"From"`    // 从哪里来
	To       []string `bson:"To"`      // 发给谁
	Subject  string   `bson:"Subject"` // 标题是啥
	SendData struct {
		Password     string
		SysTime      string // 系统时间
		Source       string // 来源
		SecurityCode string // 安全码
	} `bson:"SendData"` // 要发送邮件的数据内容
}
