package mFile

import (
	"path"
)

type CompressImgOpt struct {
	SavePath string // 新图片路径
	Replace  bool
	Src      string // 图片地址
	Email    string
	ApiKey   string
}

func CompressImg(opt CompressImgOpt) (resData string, resErr error) {
	resData, err := Tinypng(TinyOpt{
		Src:    opt.Src,
		Email:  opt.Email,
		ApiKey: opt.ApiKey,
	})
	if err != nil {
		resErr = err
		return
	}

	SavePath := opt.SavePath
	if opt.Replace {
		SavePath = opt.Src
	}

	SaveName := path.Base(SavePath)

	LocalRes, err := DownFile(DownFileOpt{
		Url:      resData,
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
