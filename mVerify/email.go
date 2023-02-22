package mVerify

import "regexp"

func IsEmail(str string) bool {
	// pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}

// 1-12 位字母数字下划线和中文
func IsNickName(str string) bool {
	pattern := "^[a-zA-Z0-9_\u4e00-\u9fa5]{1,12}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}
