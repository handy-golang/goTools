package mFile

import (
	"os"
	"path/filepath"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mTime"

	"github.com/gocolly/colly"
)

type SaveParam struct {
	Url       string // 文件Url
	Name      string // 保存的文件名，会自动添加后缀名
	Suffix    string // 文件后缀名
	Directory string // 保存的目录，可以为相对路径
}

// 下载文件并保存
func SaveRemote(param SaveParam) {
	directory := param.Directory
	if len(directory) < 1 {
		directory = "./"
	}

	Name := param.Name
	if len(Name) < 1 {
		Name = mTime.GetUnix()
	}

	Url := param.Url
	if len(Url) < 5 {
		Url = "http://bz.mo7.cc/api/public/bz?idx=1"
	}

	lastName := param.Suffix
	if len(lastName) < 3 {
		lastName = filepath.Ext(Url)
	}

	path, _ := filepath.Abs(directory)

	fileName := Name + lastName

	isPath := mPath.Exists(path)

	if !isPath {
		os.Mkdir(path, os.ModePerm)
	}

	filePath := filepath.Clean(path + string(os.PathSeparator) + fileName)

	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		file, _ := os.Create(filePath)
		file.Write(r.Body)
	})
	c.Visit(param.Url)
}
