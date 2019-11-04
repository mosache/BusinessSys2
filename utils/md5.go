package utils

import (
	"crypto/md5"
	"fmt"
)

//生成MD5
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}
