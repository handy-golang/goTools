package mPath

import "os"

// 判断目录是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(nil) {
		return false, nil
	}
	return false, nil
}
