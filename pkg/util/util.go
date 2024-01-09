package util

import (
	"os"
	"regexp"
)

// 判断目录或者文件不存在
func IsNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return true, nil
	}

	return false, err
}

// 判断是否是手机号
func IsMobile(mobile string) bool {
	ok, _ := regexp.MatchString(`^1[3456789]{1}\d{9}$`, mobile)

	return ok
}

// 判断账号和密码格式
func IsAccountAndPassword(text string) bool {
	ok, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, text)

	return ok
}
