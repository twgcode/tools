/**
@Author: wei-g
@Date:   2020/9/3 11:53 上午
@Description:
*/

package simplefilepath

import (
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
