package mJson

import (
	"bytes"
	"encoding/json"

	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

func ToJson(data any) []byte {
	jsonStr, err := jsoniter.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return jsonStr
}

func ToStr(data any) string {
	byteStr := ToJson(data)

	jsonStr := mStr.ToStr(byteStr)

	return jsonStr
}

func JsonFormat(jsonByte []byte) string {
	var out bytes.Buffer
	json.Indent(&out, jsonByte, "", "	")

	return out.String()
}

func Format(data any) string {
	jsonByte := ToJson(data)
	return JsonFormat(jsonByte)
}
