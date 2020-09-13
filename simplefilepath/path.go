/**
@Author: wei-g
@Date:   2020/9/3 11:53 上午
@Description:
*/

package simplefilepath

import (
	"os"
	"runtime"
	"time"
)

var pathTimeLayout = "2006/01/02"

func init() {
	if runtime.GOOS == "windows" {
		pathTimeLayout = `2006\01\02`
	}
}

// NowDayTimePath 根据当前时间的年月日，生成一个时间的相对路径路径
func NowDayTimePath() (path string) {
	path = time.Now().Format(pathTimeLayout)
	return
}

// PathIsExist 判断 文件或者目录是存在
func PathIsExist(path string) (ok bool) {
	var (
		err error
	)
	_, err = os.Stat(path)
	return err == nil || os.IsExist(err)
}
