package mTikker

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/tmpl"
)

func (obj *TikkerObj) Run() {
	fileName := mStr.Join(
		"t_", mEncrypt.RandStr(3), ".sh",
	)

	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(tmpl.TikkerSh))
	Tmpl.Execute(Body, tmpl.TikkerShParam{
		Path:      obj.Path,
		FileName:  fileName,
		ShellCont: obj.Shell,
		LogPath:   obj.LogPath,
	})

	Cont := Body.String()
	filePath := mStr.Join(
		obj.Path,
		mStr.ToStr(os.PathSeparator),
		fileName,
	)

	mFile.Write(filePath, Cont)

	_, err := exec.Command("pm2", "start", filePath, "--name", fileName, "--no-autorestart").Output()
	if err != nil {
		obj.Log.Println("执行失败", mStr.ToStr(err))
	} else {
		obj.Log.Println("执行成功")
	}
}
