package global

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mCycle"
)

func Start() {
	// 初始化目录列表
	config.DirInit()

	// 初始化日志系统 保证日志可用
	mCycle.New(mCycle.Opt{
		Func:      LogInt,
		SleepTime: time.Hour * 8,
	}).Start()

	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(AppInfo))
	Tmpl.Execute(Body, AppInfoParam{
		Version: config.AppInfo.Version,
	})
	Cont := Body.String()

	config.LoadAppEnv()

	Log.Println(Cont)
	fmt.Println(Cont)
}
