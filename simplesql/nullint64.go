/**
@Author: wei-g
@Date:   2020/12/17 3:07 下午
@Description:
*/

package simplesql

import (
	"database/sql"
	"encoding/json"
	"strconv"
)

// NullInt64 封装 sql.NullInt64 类型
type NullInt64 struct {
	sql.NullInt64
}

// NewNullInt64 构造函数
func NewNullInt64(Int64 int64, Valid bool) *NullInt64 {
	return &NullInt64{sql.NullInt64{
		Int64: Int64,
		Valid: Valid,
	}}

}

// UnmarshalJSON json 反序列化接口
func (n *NullInt64) UnmarshalJSON(b []byte) (err error) {
	var (
		s     string
		value int64
	)
	if string(b) == `null` {
		n.Valid = false
		n.Int64 = 0
		return
	}

	// Try string first
	if err = json.Unmarshal(b, &s); err != nil {
		// Fallback to number
		if err = json.Unmarshal(b, &n.Int64); err != nil {
			return
		}
		n.Valid = true
		return nil
	}
	if value, err = strconv.ParseInt(s, 10, 64); err != nil {
		return
	}
	n.Int64 = value
	n.Valid = true
	return

}

// MarshalJSON json 序列化接口
func (n NullInt64) MarshalJSON() (data []byte, err error) {
	if n.Valid {
		return json.Marshal(n.Int64)
	} else {
		return json.Marshal(nil)
	}

}
