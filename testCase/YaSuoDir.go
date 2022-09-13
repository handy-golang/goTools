package testCase

import (
	"os"
	"path"
	"strings"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
	"github.com/EasyGolang/goTools/mJson"
)

var FileAllPath = []string{}

type TinyKey struct {
	Email  string
	ApiKey string
	Index  int
}

var YaSuoDirPath = "/root/ProdProject/file.unido-itpo-beijing.cn"

var TinyList = map[string]TinyKey{
	"tiny2@mo7.cc": {
		Email:  "tiny2@mo7.cc",
		ApiKey: "dctdH18DxC4KZzPtwKKqJWMnW9s2Kk6m",
		Index:  1,
	},
	"tiny3@mo7.cc": {
		Email:  "tiny3@mo7.cc",
		ApiKey: "pxyZR1dDJdwnmtBdJxCQkdDQ7KnXC8rF",
		Index:  1,
	},
	"tiny4@mo7.cc": {
		Email:  "tiny4@mo7.cc",
		ApiKey: "MSYJfgBHLZ98KzzxGz6LS9NH6lNNXDM2",
		Index:  1,
	},
	"tiny5@mo7.cc": {
		Email:  "tiny5@mo7.cc",
		ApiKey: "lWZ9ZSVtL0ss2cw0s4gnJLQ4sgShqRNY",
		Index:  1,
	},
	"tiny6@mo7.cc": {
		Email:  "tiny6@mo7.cc",
		ApiKey: "YkS7RTpgnH7ZxWvPt7rSHCtBg35BvzVM",
		Index:  1,
	},
	"122198384@qq.com": {
		Email:  "tiny2@qq.com",
		ApiKey: "nNCj8h9ffgJHkDjJwrpqVTffVcsrnr6d",
		Index:  20,
	},
}

func YaSuoDir() {
	EachDir(YaSuoDirPath)
}

func EachDir(T string) {
	DirList, _ := os.ReadDir(T)
	for _, DirObj := range DirList {
		name := DirObj.Name()
		if DirObj.IsDir() {
			path := T + "/" + name
			EachDir(path)
		} else {
			File := T + "/" + name

			extName := path.Ext(File)
			extName = strings.ToLower(extName)

			if extName == ".jpeg" || extName == ".png" || extName == ".webp" || extName == ".jpg" {
				YaSuoFunc(File)
			}
		}
	}
}

func YaSuoFunc(path string) {
	SelectKey := TinyKey{}
	for _, TinyKey := range TinyList {
		if TinyKey.Index < 480 {
			SelectKey = TinyKey
			break
		}
	}
	global.Log.Println("开始压缩", SelectKey.Email, SelectKey.Index)
	resData, err := mFile.CompressImg(mFile.CompressImgOpt{
		Replace: false,
		Src:     path,
		Email:   SelectKey.Email,
		ApiKey:  SelectKey.ApiKey,
	})
	if err != nil {
		global.LogErr("压缩失败", path, SelectKey.Email)
		return
	}
	SelectKey.Index += 1
	TinyList[SelectKey.Email] = SelectKey
	global.Log.Println("压缩结束", resData)
	mFile.Write(config.Dir.JsonData+"/TinyList.json", mJson.ToStr(TinyList))
}
