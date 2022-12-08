package mEmail

import (
	"bytes"
	_ "embed"
	"net/smtp"
	"text/template"

	"github.com/EasyGolang/goTools/mStr"
	"github.com/jordan-wright/email"
)

// email-template.html  文件内容
/*
<body>
    <div>{{.Message}}</div>
    <div>{{.SysTime}}</div>
    <div>
        github.com/EasyGolang/goTools
    </div>
</body>
*/

// TStr 变量
/*

//go:embed email-template.html
var TStr string

*/

//使用方式
/*

err := mEmail.New(mEmail.Opt{
	Account:  "hunter_data_center@mo7.cc",
	Password: "hIXY2pYSuxEz6Y5k",
	To: []string{
		"meichangliang@mo7.cc",
	},
	From:        "Hunter服务",
	Subject:     "这里是Subject",
	Port:        "587",
	Host:        "smtp.feishu.cn",
	TemplateStr: TStr,
	SendData: struct {
		SysTime string
		Message string
	}{
		SysTime: time.Now().Format("2006-01-02 15:04:05"),
		Message: "这里是一封邮件",
	},
}).Send()





*/

type Opt struct {
	Account     string
	Password    string
	To          []string
	From        string
	Subject     string
	Port        string
	Host        string
	TemplateStr string
	SendData    any
}

type EmailInfo struct {
	Account  string
	Password string
	From     string
	To       []string
	Subject  string
	Template *template.Template
	SendData any
	Host     string
	Port     string
}

func New(opt Opt) *EmailInfo {
	var NewOpt EmailInfo

	tel, err := template.New("").Parse(opt.TemplateStr)
	if err != nil {
		// errStr := fmt.Errorf("模板字符构建失败")
		return nil
	}

	NewOpt.Account = opt.Account
	NewOpt.Password = opt.Password
	NewOpt.From = opt.From
	NewOpt.To = opt.To
	NewOpt.Subject = opt.Subject
	NewOpt.Template = tel
	NewOpt.SendData = opt.SendData
	NewOpt.Host = opt.Host
	NewOpt.Port = opt.Port

	return &NewOpt
}

func (Info *EmailInfo) Send() error {
	em := email.NewEmail()

	em.From = mStr.Join(
		Info.From, "<", Info.Account, ">",
	)
	em.To = Info.To
	em.Subject = Info.Subject

	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	Info.Template.Execute(body, Info.SendData)
	// html形式的消息
	em.HTML = body.Bytes()

	addr := mStr.Join(
		Info.Host, ":", Info.Port,
	)
	err := em.Send(addr, smtp.PlainAuth("", Info.Account, Info.Password, Info.Host))

	return err
}
