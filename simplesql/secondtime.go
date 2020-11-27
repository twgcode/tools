/**
@Author: wei-g
@Date:   2020/11/26 6:11 下午
@Description: sql.NullTime 时间类型对json序列化时进行封装
*/

package simplesql

import (
	"database/sql"
	"github.com/twgcode/tools/simpletime"
	"time"
)

// SecondNullTime 精确到秒
type SecondNullTime struct {
	sql.NullTime
}

// ToNullTime 根据 time类型快速生成 SecondNullTime
func ToSecondNullTime(t time.Time) SecondNullTime {
	return SecondNullTime{sql.NullTime{Time: t, Valid: !t.IsZero()}}
}

// UnmarshalJSON json 反序列化接口
func (s *SecondNullTime) UnmarshalJSON(b []byte) (err error) {
	if string(b) == `null` {
		s.Valid = false
		s.Time = time.Time{}
		return
	}
	t := simpletime.SecondTime(s.Time)
	if err = t.UnmarshalJSON(b); err != nil {
		s.Valid = false
		return
	}
	s.Time = time.Time(t)
	// 判断是否为零时
	if s.Time.IsZero() {
		s.Valid = false
	} else {
		s.Valid = true
	}
	return
}

// MarshalJSON json 序列化接口
func (s *SecondNullTime) MarshalJSON() (data []byte, err error) {
	if !s.Valid || s.Time.IsZero() {
		data = []byte(`null`)
		return
	}
	t := simpletime.SecondTime(s.Time)
	return t.MarshalJSON()
}

// ToNullTime 根据 time 类型快速生成 sql.NullTime
func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}
