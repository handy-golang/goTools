package mJson

import (
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
