package mJson

import (
	"log"
)

// struct 进行 json 格式 打印
func Println(data any) string {
	json := ToJson(data)
	jsonStr := JsonFormat(json)
	log.Println(jsonStr)
	return jsonStr
}
