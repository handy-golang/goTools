package mStr

import (
	"fmt"
	"os"
	"strings"
)

/*
拼接字符串
*/

/*
	a := []rune("mo7欢迎你")
	a := []byte("mo7欢迎你")
	a := 10.97
	a := os.PathSeparator

	str := mStr.ToStr(a)

*/
func ToStr(p any) string {
	// fmt.Println("type: ", reflect.TypeOf(p))
	returnStr := ""
	switch p := p.(type) {
	case []int32:
		returnStr = string(p)
	case []uint8:
		returnStr = string(p)
	case int32:
		returnStr = string(p)
	case uint8:
		returnStr = string(p)
	default:
		returnStr = fmt.Sprintf("%+v", p)
	}

	return returnStr
}

func Join(s ...string) string {
	var build strings.Builder
	for _, v := range s {
		build.WriteString(v)
	}
	return build.String()
}

/*

var config = `
app.name = ${appName}
app.ip = ${appIP}
app.port = ${appPort}
`

var dev = map[string]string{
	"appName": "my_ap123p",
	"appIP":   "0.0.0.0",
	"appPort": "8080",
}

	s := mStr.Temp(config, dev)

*/

func Temp(config string, lMap map[string]string) string {
	s := os.Expand(config, func(k string) string {
		return lMap[k]
	})
	return s
}
