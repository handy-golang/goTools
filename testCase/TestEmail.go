package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mTime"
)

func TestEmail() {
	TStr := "测试邮件" + mEncrypt.GetUUID()

	start := mTime.GetUnixInt64()
	fmt.Println("开始发送邮件", start)

	err := mEmail.New(mEmail.Opt{
		// Account:  "trade@mo7.cc",
		// Password: "svxxaeJgxxMraNmxxdVFJu",
		// Port:     "587",
		// Host:     "smtp.exmail.qq.com",

		Account:  "670188307@qq.com",
		Password: "momxxwdnjxxstnuxxxibcej",
		Port:     "587",
		Host:     "smtp.qq.com",

		To: []string{
			"meichangliang@outlook.com",
			"670188307@qq.com",
			"mo7@mo7.cc",
		},
		From:    "Hunter 测试 服务",
		Subject: "这里是Subject",

		TemplateStr: TStr,
		SendData: struct {
			SysTime string
			Message string
		}{
			SysTime: mTime.UnixFormat(mTime.GetUnix()),
			Message: "这里是一封邮件",
		},
	}).Send()

	fmt.Println("发送邮件结束", err)

	end := mTime.GetUnixInt64()

	fmt.Println("end", end)
	fmt.Println("diff", end-start)
}
