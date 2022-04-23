package mGin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/gin-gonic/gin"
)

type SPAOpt struct {
	Root         string      // 静态文件目录
	RelativePath string      // 路由
	Router       *gin.Engine // 创建好的 Gin
}

func SPAServer(opt SPAOpt) *gin.Engine {
	if len(opt.Root) < 1 {
		errStr := fmt.Errorf("缺少 Path 参数 ")
		panic(errStr)
	}

	uri := opt.Root
	router := opt.Router

	if router == nil {
		errStr := fmt.Errorf("缺少 Router 参数 ")
		panic(errStr)
	}

	router.Use(Public)

	router.StaticFile(opt.RelativePath, uri+"/index.html")

	fileInfoList, err := ioutil.ReadDir(uri)
	if err != nil {
		errorsStr := fmt.Errorf("目录读取失败")
		panic(errorsStr)
	}

	for i := range fileInfoList {
		name := fileInfoList[i].Name()
		path := uri + "/" + name

		if mPath.IsDir(path) {
			router.StaticFS(opt.RelativePath+name, http.Dir(path))
		}
		if mPath.IsFile(path) {
			router.StaticFile(opt.RelativePath+name, path)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.File(uri + "/index.html")
	})

	return router
}
