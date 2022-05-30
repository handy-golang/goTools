package mTikker

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/tmpl"
)

// pm2 安装
func (obj *TikkerObj) InstPm2() *TikkerObj {
	filePath := mPath.Dir.App
	fileName := mStr.Join(
		"i_", mEncrypt.RandStr(5), ".sh",
	)

	Body := new(bytes.Buffer)
	Tmpl := template.Must(template.New("").Parse(tmpl.InstPm2))
	Tmpl.Execute(Body, tmpl.InstPm2Param{
		Path:     filePath,
		FileName: fileName,
	})
	Cont := Body.String()

	shellPath := mStr.Join(
		filePath,
		mStr.ToStr(os.PathSeparator),
		fileName,
	)

	mFile.Write(shellPath, Cont)

	res, err := exec.Command("/bin/bash", shellPath).Output()
	if err != nil {
		obj.Log.Println("环境安装失败", mStr.ToStr(err))
	} else {
		obj.Log.Println("环境安装成功", mStr.ToStr(res))
	}

	return obj
}
