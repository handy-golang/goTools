package mGin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/gin-gonic/gin"
)

type SPAServerOpt struct {
	Path   string      // 静态文件目录
	Router *gin.Engine // 创建好的 Gin
}

func SPAServer(opt SPAServerOpt) *gin.Engine {
	uri := "./"
	router := opt.Router

	if router == nil {
		errStr := fmt.Errorf("缺少 Router 参数 ")
		panic(errStr)
	}

	router.Use(Public)

	router.StaticFile("/", uri+"/index.html")

	fileInfoList, err := ioutil.ReadDir(uri)
	if err != nil {
		errorsStr := fmt.Errorf("目录读取失败")
		panic(errorsStr)
	}

	for i := range fileInfoList {
		name := fileInfoList[i].Name()
		path := uri + "/" + name

		if mPath.IsDir(path) {
			router.StaticFS("/"+name, http.Dir(path))
		}
		if mPath.IsFile(path) {
			router.StaticFile("/"+name, path)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.File(uri + "/index.html")
	})

	return router
}
