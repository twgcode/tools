/**
@Author: wei-g
@Date:   2021/7/19 9:43 上午
@Description:
*/

package simplestruct

import (
	"encoding/json"
	"strconv"
)

type Uint64 uint64

// MarshalJSON 满足json 序列化接口
func (i Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(i), 10))
}

// UnmarshalJSON 满足json 反序列化接口
func (i *Uint64) UnmarshalJSON(b []byte) (err error) {
	var (
		s     string
		value uint64
	)
	// Try string first
	if err := json.Unmarshal(b, &s); err != nil {
		// Fallback to number
		return json.Unmarshal(b, (*uint64)(i))
	}

	if value, err = strconv.ParseUint(s, 10, 64); err != nil {
		return
	}
	*i = Uint64(value)
	return
}

// ToBuiltinUint64Slice Uint64类型切片转换为 系统内置标准的uint64类型切片
func ToBuiltinUint64Slice(slice []Uint64) (int64Slice []uint64) {
	if slice == nil {
		return nil
	}
	var tmp []Uint64

	if len(slice) <= 1024 { // 判断数据是否过大 如果过大则不进行数据copy
		// 拷贝数据, 防止其他协程进行修改
		tmp = make([]Uint64, len(slice), cap(slice))
		copy(tmp, slice)
	} else {
		tmp = slice
	}

	int64Slice = make([]uint64, 0, cap(tmp))
	// 进行数据转换
	for _, v := range tmp {
		int64Slice = append(int64Slice, uint64(v))
	}
	return
}

// ToUint64Slice  系统内置标准的uint64类型切片 转换为Uint64类型切片
func ToUint64Slice(uint64Slice []uint64) (slice []Uint64) {
	if uint64Slice == nil {
		return nil
	}
	var tmp []uint64
	if len(uint64Slice) <= 1024 { // 判断数据是否过大 如果过大则不进行数据copy
		// 拷贝数据, 防止其他协程进行修改
		tmp = make([]uint64, len(uint64Slice), cap(uint64Slice))
		copy(tmp, uint64Slice)
	} else {
		tmp = uint64Slice
	}

	slice = make([]Uint64, 0, cap(tmp))
	for _, v := range tmp {
		slice = append(slice, Uint64(v))
	}
	return
}

// Uint64Contains 判断 item是发在指定的切片中; 新版不建议使用,废弃状态
func Uint64Contains(slice []uint64, val uint64) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// AppendBuiltinUint64Slice 合并(追加方式)2个切片
func AppendBuiltinUint64Slice(s1, s2 []uint64) (slice []uint64) {
	if s1 == nil && s2 == nil {
		return nil
	}
	slice = make([]uint64, 0, len(s1)+len(s2))
	slice = append(slice, s1...)
	slice = append(slice, s2...)
	return
}

// AppendUint64Slice 合并(追加方式)2个切片
func AppendUint64Slice(s1, s2 []Uint64) (slice []Uint64) {
	if s1 == nil && s2 == nil {
		return nil
	}
	slice = make([]Uint64, 0, len(s1)+len(s2))
	slice = append(slice, s1...)
	slice = append(slice, s2...)
	return
}

// BuiltinUint64SliceContains 判断 item是发在指定的切片中
func BuiltinUint64SliceContains(slice []uint64, val uint64) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// BuiltinUint64SliceRemoveDuplicateValues 对切片进行去重
func BuiltinUint64SliceRemoveDuplicateValues(slice []uint64) []uint64 {
	if len(slice) == 0 {
		return slice
	}
	list := make([]uint64, 0, cap(slice))
	for _, item := range slice {
		if !BuiltinUint64SliceContains(list, item) {
			list = append(list, item)
		}
	}
	return list
}

// BuiltinUint64SliceRemoveElement 从切片中移除遇到第一个指定元素
func BuiltinUint64SliceRemoveElement(source []uint64, element uint64) ([]uint64, bool) {
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
func MaxBuiltinUint64(l ...int64) int64 {
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

func MinBuiltinUint64(l ...int64) int64 {
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
