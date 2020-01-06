package utils

import (
	"os"
	"strings"
)

//PathExists 判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
//IsExistsInGoPath 在GoPath中判断文件是否存在
func IsExistsInGoPath(file string) bool {
	gopaths := strings.Split(os.Getenv("GOPATH"), ";")

	for _, gp := range gopaths {
		b, err := PathExists(gp + "/src" + file)
		if err == nil && b == true {
			return true
		}
	}
	return false
}
