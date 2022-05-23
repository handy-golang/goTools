package mFile

import (
	"os"
	"path"
	"strings"

	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

type GetNameOpt struct {
	FileName string
	SavePath string
}

type GetNameType struct {
	Count    string
	Name     string
	SrcName  string
	ExtName  string
	SavePath string
}

func GetName(opt GetNameOpt) string {
	extName := path.Ext(opt.FileName)                      // 后缀名
	name := strings.Replace(opt.FileName, extName, "", -1) // 把后缀名换成空字符串

	var Obj GetNameType
	Obj.Count = "0"
	Obj.Name = name
	Obj.SrcName = name
	Obj.ExtName = extName
	Obj.SavePath = opt.SavePath

	isThere := Obj.FileThere()

	if isThere {
		Obj.GetNewName()
	}

	return mStr.Join(
		Obj.Name, Obj.ExtName,
	)
}

func (obj *GetNameType) FileThere() bool {
	filePath := mStr.Join(
		obj.SavePath,
		mStr.ToStr(os.PathSeparator),
		obj.Name,
		obj.ExtName,
	)

	isFilePath := mPath.Exists(filePath)

	return isFilePath
}

func (obj *GetNameType) GetNewName() *GetNameType {
	obj.Count = mCount.Add(obj.Count, "1")
	obj.Name = mStr.Join(
		obj.SrcName, "_", obj.Count,
	)

	isThere := obj.FileThere()
	if isThere {
		obj.GetNewName()
	}

	return obj
}
