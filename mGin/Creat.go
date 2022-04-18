package mGin

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/gin-gonic/gin"
)

func SPAServer(uri string, router *gin.Engine) {
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
}
