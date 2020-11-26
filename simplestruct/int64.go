/**
@Author: wei-g
@Date:   2020/11/13 10:02 上午
@Description: 一些基本的数据类型封装
*/

package simplestruct

import (
	"encoding/json"
	"strconv"
)

// Int64Str
// 实现主要来源于 https://stackoverflow.com/questions/49415573/golang-json-how-do-i-unmarshal-array-of-strings-into-int64?rq=1
type Int64 int64

// MarshalJSON 满足json 序列化接口
func (i Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(i), 10))
}

// UnmarshalJSON 满足json 反序列化接口
func (i *Int64) UnmarshalJSON(b []byte) (err error) {
	var (
		s     string
		value int64
	)
	// Try string first
	if err := json.Unmarshal(b, &s); err != nil {
		// Fallback to number
		return json.Unmarshal(b, (*int64)(i))
	}
	if value, err = strconv.ParseInt(s, 10, 64); err != nil {
		return
	}
	*i = Int64(value)
	return
}

// ToBuiltinInt64Slice Int64类型切片转换为 系统内置标准的int64类型切片
func ToBuiltinInt64Slice(slice []Int64) (int64Slice []int64) {
	if slice == nil {
		return nil
	}
	var tmp []Int64

	if len(slice) <= 1024 { // 判断数据是否过大 如果过大则不进行数据copy
		// 拷贝数据, 防止其他协程进行修改
		tmp = make([]Int64, len(slice), cap(slice))
		copy(tmp, slice)
	} else {
		tmp = slice
	}

	int64Slice = make([]int64, 0, cap(tmp))
	// 进行数据转换
	for _, v := range tmp {
		int64Slice = append(int64Slice, int64(v))
	}
	return
}

// ToInt64Slice  系统内置标准的int64类型切片 转换为Int64类型切片
func ToInt64Slice(int64Slice []int64) (slice []Int64) {
	if int64Slice == nil {
		return nil
	}
	var tmp []int64
	if len(int64Slice) <= 1024 { // 判断数据是否过大 如果过大则不进行数据copy
		// 拷贝数据, 防止其他协程进行修改
		tmp = make([]int64, len(int64Slice), cap(int64Slice))
		copy(tmp, int64Slice)
	} else {
		tmp = int64Slice
	}

	slice = make([]Int64, 0, cap(tmp))
	for _, v := range tmp {
		slice = append(slice, Int64(v))
	}
	return

}
