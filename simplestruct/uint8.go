/**
@Author: wei-g
@Date:   2021/9/23 11:02 上午
@Description:
*/

package simplestruct

// BuiltinUint8SliceContains 判断 item是发在指定的切片中
func BuiltinUint8SliceContains(slice []uint8, val uint8) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// BuiltinUint8SliceDuplicateValue 切片中是否有重复的值
func BuiltinUint8SliceDuplicateValue(strSlice []uint8) bool {
	if len(strSlice) == 0 {
		return false
	}
	_map := make(map[uint8]struct{}, len(strSlice))
	for _, v := range strSlice {
		if _, ok := _map[v]; ok {
			return true
		}
		_map[v] = struct{}{}
	}
	return false
}

// BuiltinUint8SliceRemoveDuplicateValues 对切片进行去重
func BuiltinUint8SliceRemoveDuplicateValues(slice []uint8) []uint8 {
	if len(slice) == 0 {
		return slice
	}
	list := make([]uint8, 0, cap(slice))
	for _, item := range slice {
		if !BuiltinUint8SliceContains(list, item) {
			list = append(list, item)
		}
	}
	return list
}

// BuiltinUint8SliceRemoveElement 从切片中移除遇到第一个指定元素
func BuiltinUint8SliceRemoveElement(source []uint8, element uint8) ([]uint8, bool) {
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

func MaxBuiltinUint8(l ...int8) int8 {
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

func MinBuiltinUint8(l ...int8) int8 {
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
