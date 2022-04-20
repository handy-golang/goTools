package mJson

import jsoniter "github.com/json-iterator/go"

func ToJson(data any) []byte {
	jsonStr, err := jsoniter.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return jsonStr
}
