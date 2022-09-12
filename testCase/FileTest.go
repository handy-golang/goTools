package testCase

import (
	"fmt"
	"log"

	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mFile"
)

func FileTest() {
	// TinyFetch()
	// DownFile()

	resData, err := mFile.CompressImg(mFile.CompressImgOpt{
		Replace: true,
		Src:     config.Dir.App + "/jsonData/10.png",
		Email:   "meichangliang@mo7.cc",
		ApiKey:  "Hl6wpxNdBg0Dvv2s7BcVsKks1tFZ2wBl",
	})

	fmt.Println("压缩图片", resData, err)
}

func TinyFetch() {
	log.Println("开始压缩")
	resData, err := mFile.Tinypng(mFile.TinyOpt{
		Src:    config.Dir.App + "/jsonData/2.png",
		Email:  "meichangliang@mo7.cc",
		ApiKey: "Hl6wpxNdBg0Dvv2s7BcVsKks1tFZ2wBl",
	})
	if err != nil {
		fmt.Println("err", err)
	}
	log.Println("获得压缩链接", resData)
}

func DownFile() {
	url := "https://api.tinify.com/output/dk417xcw5q5c34hm96puagmxn62t4zpd"

	resData, err := mFile.DownFile(mFile.DownFileOpt{
		Url:      url,
		SavePath: "./jsonData",
		SaveName: "2.png",
	})

	fmt.Println(resData, err)
}
