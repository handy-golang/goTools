package mJson

import "fmt"

func Println(data any) string {
	json := ToJson(data)
	jsonStr := JsonFormat(json)
	fmt.Println(jsonStr)
	return jsonStr
}
