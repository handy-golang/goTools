package mLog

import (
	"log"
	"os"
	"time"

	"github.com/EasyGolang/goTools/mPath"
)

type NewLogParam struct {
	Path string
	Name string
}

func NewLog(param NewLogParam) *log.Logger {
	FilePath := param.Path
	LogName := param.Name

	if len(FilePath) < 1 {
		FilePath = "./logs"
	}

	// 检测 logs 目录
	isLogPath := mPath.Exists(FilePath)
	if !isLogPath {
		// 不存在则创建 logs 目录
		os.Mkdir(FilePath, 0o777)
	}

	file := FilePath + "/" + LogName + "-" + time.Now().Format("06年1月02日15时") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o777)
	if nil != err {
		panic(err)
	}
	loger := log.New(logFile, LogName+"-", log.Ldate|log.Ltime|log.Lshortfile)
	loger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// loger.Println("dsfsdfsdf")
	// loger.Output(2, "打印一条日志信息")
	// loger.Printf("第%d行 内容:%s", 11, "我是错误k")
	// loger.Fatal("我是错误1")
	// loger.Panic("我是错误5")
	// loger.Printf("第%d行 内容:%s", 22, "我是错误")

	return loger
}
