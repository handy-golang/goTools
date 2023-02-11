package mTask

// 邮件任务需要的基本
type SendEmail struct {
	From     string         `bson:"From"`     // 从哪里来
	To       []string       `bson:"To"`       // 发给谁
	Subject  string         `bson:"Subject"`  // 标题是啥
	Template string         `bson:"Template"` // 选择哪个邮件模板来解析
	SendData map[string]any `bson:"SendData"` // 要发送邮件的数据内容
}

// 邮件解析模板
type SysEmail struct {
	Title        string // 标题
	Message      string // 消息
	Content      string // 内容
	SysTime      string // 系统时间
	Source       string // 来源
	SecurityCode string // 安全码
}
type CodeEmail struct {
	VerifyCode   string
	Action       string
	Minute       string // 分钟数
	SysTime      string // 系统时间
	Source       string // 来源
	SecurityCode string // 安全码
}
type RegisterSucceedEmail struct {
	Password     string
	SysTime      string // 系统时间
	Source       string // 来源
	SecurityCode string // 安全码
}
