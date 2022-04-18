package mStr

import (
	"os"
	"strings"
)

/*
拼接字符串
*/

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
