package mGin

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Public(c *gin.Context) {
	fileHeader(c)

	c.Next()
}

func fileHeader(c *gin.Context) {
	c.Writer.Header().Del("Server-Type")
	c.Header("Server-Type", "goTools-SPA-Server")

	path := c.FullPath()

	fileSuffix := filepath.Ext(path)
	if fileSuffix == ".js" {
		c.Header("Content-Type", "text/javascript; charset=utf-8")
	}

	if fileSuffix == ".css" {
		c.Header("Content-Type", "text/css; charset=utf-8")
	}

	if path == "/" || fileSuffix == ".html" {
		c.Header("Content-Type", "text/html; charset=utf-8")
	}
}
