/**
@Author: wei-g
@Date:   2021/9/23 11:02 上午
@Description:
*/

package simplestruct

// BuiltinInt8SliceContains 判断 item是发在指定的切片中
func BuiltinInt8SliceContains(slice []int8, val int8) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// BuiltinInt8SliceDuplicateValue 切片中是否有重复的值
func BuiltinInt8SliceDuplicateValue(strSlice []int8) bool {
	if len(strSlice) == 0 {
		return false
	}
	_map := make(map[int8]struct{}, len(strSlice))
	for _, v := range strSlice {
		if _, ok := _map[v]; ok {
			return true
		}
		_map[v] = struct{}{}
	}
	return false
}

// BuiltinInt8SliceRemoveDuplicateValues 对切片进行去重
func BuiltinInt8SliceRemoveDuplicateValues(slice []int8) []int8 {
	if len(slice) == 0 {
		return slice
	}
	list := make([]int8, 0, cap(slice))
	for _, item := range slice {
		if !BuiltinInt8SliceContains(list, item) {
			list = append(list, item)
		}
	}
	return list
}

// BuiltinInt8SliceRemoveElement 从切片中移除遇到第一个指定元素
func BuiltinInt8SliceRemoveElement(source []int8, element int8) ([]int8, bool) {
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

func MaxBuiltinInt8(l ...int8) int8 {
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

func MinBuiltinInt8(l ...int8) int8 {
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
