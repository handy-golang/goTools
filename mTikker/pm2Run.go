package mTikker

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/tmpl"
)

func (obj *TikkerObj) RunToPm2() error {
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

	res, err := exec.Command("pm2", "start", filePath, "--name", fileName, "--no-autorestart").Output()
	if err != nil {
		errStr := fmt.Errorf("执行失败:%+v", mStr.ToStr(err))
		obj.Log.Println(errStr)
		return errStr
	} else {
		obj.Log.Println("执行成功", mStr.ToStr(res))
		return nil
	}
}

func (obj *TikkerObj) RunToShell() error {
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
	shellPath := mStr.Join(
		obj.Path,
		mStr.ToStr(os.PathSeparator),
		fileName,
	)

	mFile.Write(shellPath, Cont)

	res, err := exec.Command("/bin/bash", shellPath).Output()
	if err != nil {
		errStr := fmt.Errorf("执行失败:%+v", mStr.ToStr(err))
		obj.Log.Println(errStr)
		return errStr
	} else {
		obj.Log.Println("执行成功", mStr.ToStr(res))
		return nil
	}
}
