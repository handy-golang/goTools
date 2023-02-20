package mJson

import (
	"bytes"
	"encoding/json"

	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

// struct 转成 byte
func ToJson(data any) []byte {
	jsonStr, err := jsoniter.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return jsonStr
}

// struct 转成  string
func ToStr(data any) string {
	byteStr := ToJson(data)
	jsonStr := mStr.ToStr(byteStr)
	return jsonStr
}

// byte 格式化 json 体
func JsonFormat(jsonByte []byte) string {
	var out bytes.Buffer
	json.Indent(&out, jsonByte, "", " ")
	return out.String()
}

// struct 格式化 json 体
func Format(data any) string {
	jsonByte := ToJson(data)
	return JsonFormat(jsonByte)
}

// struct 转成 map
func StructToMap(val any) (resData map[string]any) {
	jsonByte := ToJson(val)
	err := jsoniter.Unmarshal(jsonByte, &resData)
	if err != nil {
		resData = map[string]any{}
	}
	return
}
