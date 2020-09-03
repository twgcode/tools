/**
@Author: wei-g
@Date:   2020/4/8 9:22 下午
@Description: 实现文件加密有关
*/

package sparcrypto

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5(path string) (Md5Str string, err error) {
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()
	md5obj := md5.New()
	if _, err = io.Copy(md5obj, file); err != nil {
		return
	}
	Md5Str = hex.EncodeToString(md5obj.Sum(nil))
	return
}
