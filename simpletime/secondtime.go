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

const (
	SecondTimeLayoutNano = "2006-01-02 15:04:05 MST" // 精确到秒,带时区信息
	//TimeLayoutNano       = "2006-01-02 15:04:05.000000000 MST" // 精确到毫秒,带时区信息
)

// SecondTime 精确到秒的时间序列化
type SecondTime time.Time

// UnmarshalJSON json 反序列化接口
func (s *SecondTime) UnmarshalJSON(b []byte) (err error) {
	data := strings.Trim(string(b), `"`)
	if data == "null" {
		*s = SecondTime(time.Time{})
		return
	}
	var now time.Time
	if now, err = time.Parse(SecondTimeLayoutNano, data); err != nil {
		return
	}
	*s = SecondTime(now)
	return
}

// MarshalJSON json 序列化接口
func (s *SecondTime) MarshalJSON() (data []byte, err error) {
	if time.Time(*s).IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, time.Time(*s).Format(SecondTimeLayoutNano))), nil
}

// IsSet 判断Time是否为正确设置
func (s *SecondTime) IsSet() bool {
	return !time.Time(*s).IsZero()
}
