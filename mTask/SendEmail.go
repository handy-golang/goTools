package mTask

// 邮件任务需要的基本
type SendEmail struct {
	From     string         `bson:"From"`     // 从哪里来
	To       []string       `bson:"To"`       // 发给谁
	Subject  string         `bson:"Subject"`  // 标题是啥
	Template string         `bson:"Template"` // 选择哪个邮件模板来解析
	SendData map[string]any `bson:"SendData"` // 要发送邮件的数据内容
}
