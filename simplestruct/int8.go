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
