package mFile

import (
	"path"
	"strings"
)

type CompressImgOpt struct {
	SavePath string // 新图片保存路径
	Replace  bool
	Src      string // 图片地址
	Email    string
	ApiKey   string
}

func CompressImg(opt CompressImgOpt) (resData string, resErr error) {
	TinyRes, err := Tinypng(TinyOpt{
		Src:    opt.Src,
		Email:  opt.Email,
		ApiKey: opt.ApiKey,
	})
	if err != nil {
		resErr = err
		return
	}

	SavePath := opt.SavePath
	SaveName := path.Base(opt.Src)
	if opt.Replace {
		SavePath = strings.Replace(opt.Src, SaveName, "", 1)
	} else {
		extName := path.Ext(SaveName)                      // 后缀名
		name := strings.Replace(SaveName, extName, "", -1) // 把后缀名换成空字符串
		SaveName = name + "_z" + extName
	}

	LocalRes, err := DownFile(DownFileOpt{
		Url:      TinyRes,
		SavePath: SavePath,
		SaveName: SaveName,
	})
	if err != nil {
		resErr = err
		return
	}

	resData = LocalRes

	return
}
