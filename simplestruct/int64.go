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

// Int64
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

// IntContains 判断 item 是否存在 切片中; 新版不建议使用,废弃状态
func IntContains(slice []int64, val int64) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// AppendBuiltinInt64Slice 合并(追加方式)2个切片
func AppendBuiltinInt64Slice(s1, s2 []int64) (slice []int64) {
	if s1 == nil && s2 == nil {
		return nil
	}
	slice = make([]int64, 0, len(s1)+len(s2))
	slice = append(slice, s1...)
	slice = append(slice, s2...)
	return
}

// AppendInt64Slice 合并(追加方式)2个切片
func AppendInt64Slice(s1, s2 []Int64) (slice []Int64) {
	if s1 == nil && s2 == nil {
		return nil
	}
	slice = make([]Int64, 0, len(s1)+len(s2))
	slice = append(slice, s1...)
	slice = append(slice, s2...)
	return
}

// BuiltinInt64SliceContains 判断 item是发在指定的切片中
func BuiltinInt64SliceContains(slice []int64, val int64) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// BuiltinInt64SliceDuplicateValue 切片中是否有重复的值
func BuiltinInt64SliceDuplicateValue(strSlice []int64) bool {
	if len(strSlice) == 0 {
		return false
	}
	_map := make(map[int64]struct{}, len(strSlice))
	for _, v := range strSlice {
		if _, ok := _map[v]; ok {
			return true
		}
		_map[v] = struct{}{}
	}
	return false
}

// BuiltinInt64SliceRemoveDuplicateValues 对切片进行去重
func BuiltinInt64SliceRemoveDuplicateValues(slice []int64) []int64 {
	if len(slice) == 0 {
		return slice
	}
	list := make([]int64, 0, cap(slice))
	for _, item := range slice {
		if !BuiltinInt64SliceContains(list, item) {
			list = append(list, item)
		}
	}
	return list
}

// BuiltinInt64SliceRemoveElement 从切片中移除遇到第一个指定元素
func BuiltinInt64SliceRemoveElement(source []int64, element int64) ([]int64, bool) {
	if len(source) == 0 {
		return source, false
	}
	for i, v := range source {
		if v == element {
			source = append(source[:i], source[i+1:]...)
			return source, true
		}
	}
	return source, false
}

func MaxBuiltinInt64(l ...int64) int64 {
	if len(l) == 0 {
		return 0
	}
	if len(l) == 1 {
		return l[0]
	}
	result := l[0]
	for _, v := range l {
		if v > result {
			result = v
		}
	}
	return result
}

func MinBuiltinInt64(l ...int64) int64 {
	if len(l) == 0 {
		return 0
	}
	if len(l) == 1 {
		return l[0]
	}
	result := l[0]
	for _, v := range l {
		if v < result {
			result = v
		}
	}
	return result
}
