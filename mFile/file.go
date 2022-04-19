package mFile

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/EasyGolang/goTools/mPath"
)

// 写入文件内容 fileName 为文件的路径
func Write(fileName string, content string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		panic("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		f.WriteAt([]byte(content), n)
		log.Println(fileName, "write succeed!")
		defer f.Close()
	}
}

// 获取一个文件的类型
func GetContentType(fileName string) string {
	if !mPath.IsFile(fileName) {
		errorsStr := fmt.Errorf("fileName必须为一个文件")
		panic(errorsStr)
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic("读取文件出错")
	}
	defer file.Close()

	// Only the first 512 bytes are used to sniff the content type.
	bs := make([]byte, 512)
	_, err = file.Read(bs)
	if err != nil {
		return "application/octet-stream"
	}

	// // Use the net/http package's handy DectectContentType function. Always returns a valid
	// // content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(bs)

	return contentType
}

func ReadFile(fileName string) []byte {
	if !mPath.IsFile(fileName) {
		errStr := fmt.Errorf("fileName必须为一个文件")
		log.Println(errStr)
		return []byte("")
	}

	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte("")
	}
	return f
}
