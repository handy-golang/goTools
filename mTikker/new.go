package mTikker

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"os/exec"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/EasyGolang/goTools/tmpl"
)

type TikkerOpt struct {
	ShellContent string
	LogPath      string
}

type TikkerObj struct {
	Path    string
	Shell   string
	Log     *log.Logger
	LogPath string
}

func NewTikker(opt TikkerOpt) *TikkerObj {
	var obj TikkerObj

	// 生成脚本执行目录
	Path := mPath.Dir.App
	isPath := mPath.Exists(Path)
	if !isPath {
		// 不存在则创建目录
		os.Mkdir(Path, 0o777)
	}

	// 日志存放目录
	LogPath := mPath.Dir.App + "/logs"
	if len(opt.LogPath) > 0 {
		LogPath = opt.LogPath
	}

	// 创建日志目录
	isLogPath := mPath.Exists(LogPath)
	if !isLogPath {
		// 不存在则创建 logs 目录
		os.Mkdir(LogPath, 0o777)
	}

	obj.Path = Path
	obj.Shell = opt.ShellContent

	// 创建日志文件
	file := LogPath + "/" + "mTikker.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o777)
	if nil != err {
		panic(err)
	}
	obj.Log = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	obj.Log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	obj.LogPath = file

	obj.Log.Println("欢迎使用 goTools !")

	return &obj
}

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
