/**
@Author: wei-g
@Date:   2020/11/25 3:01 下午
@Description: 简化时间序列化操作等有关处理
*/
// 本文件部分代码实现借鉴 https://www.liwenzhou.com/posts/Go/json_tricks_in_go/#autoid-0-0-12
package simpletime

import (
	"fmt"
	"strings"
	"time"
)

var (
	// 代表 null 时间
	nilTime = (time.Time{}).UnixNano()
)

const (
	//TimeLayoutNano = "2006-01-02 15:04:05.000000000 MST" // 精确到毫秒,带时区信息
	SecondTimeLayoutNano = "2006-01-02 15:04:05 MST" // 精确到毫秒,带时区信息
)

// SecondTime 精确到秒的时间序列化
type SecondTime struct {
	time.Time
}

// UnmarshalJSON json 反序列化接口
func (t *SecondTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(SecondTimeLayoutNano, s)
	return
}

// MarshalJSON json 序列化接口
func (t *SecondTime) MarshalJSON() ([]byte, error) {
	if t.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, t.Time.Format(SecondTimeLayoutNano))), nil
}

// IsSet 判断Time是否为正确设置
func (t *SecondTime) IsSet() bool {
	return t.UnixNano() != nilTime
}
