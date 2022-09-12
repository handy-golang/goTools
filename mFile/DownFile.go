package mFile

import (
	"bytes"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/gocolly/colly"
)

type DownFileOpt struct {
	Url      string
	SavePath string
	SaveName string
}

func DownFile(opt DownFileOpt) (resData string, resErr error) {
	Url := opt.Url

	SavePath := opt.SavePath
	if len(SavePath) < 1 {
		SavePath = "."
	}
	isLogPath := mPath.Exists(SavePath)
	if !isLogPath {
		os.Mkdir(SavePath, 0o777)
	}

	SavePath, _ = filepath.Abs(SavePath)

	SaveName := opt.SaveName
	if len(SaveName) < 2 {
		SaveName = GetName(GetNameOpt{
			FileName: mEncrypt.RandStr(5),
			SavePath: SavePath,
			RandName: true,
		})
	}

	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		fileName := SaveName
		extName := path.Ext(SaveName) // 后缀名
		if len(extName) < 1 {
			extName = ContentToExtName(r.Headers.Get("Content-Type"))
			if len(extName) > 0 {
				extName = "." + extName
			}
			fileName = fileName + extName
		}

		SaveFile := SavePath + "/" + fileName
		f, err := os.Create(SaveFile)
		if err != nil {
			resErr = err
		}
		io.Copy(f, bytes.NewReader(r.Body))

		resData = SaveFile
	})
	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			resErr = err
		}
	})
	c.Visit(Url)

	return
}
