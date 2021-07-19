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

//  Uint64Contains 判断 item是发在指定的切片中
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
