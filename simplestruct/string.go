/**
@Author: wei-g
@Date:   2021/8/24 5:20 下午
@Description:
*/

package simplestruct

// BuiltinStringSliceContains 判断 item是发在指定的切片中
func BuiltinStringSliceContains(slice []string, val string) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// BuiltinStringSliceRemoveDuplicateValues 对切片进行去重
func BuiltinStringSliceRemoveDuplicateValues(strSlice []string) []string {
	if len(strSlice) == 0 {
		return strSlice
	}
	list := make([]string, 0, cap(strSlice))
	for _, item := range strSlice {
		if !BuiltinStringSliceContains(list, item) {
			list = append(list, item)
		}
	}
	return list
}

// BuiltinStringSliceDuplicateValue 切片中是否有重复的值
func BuiltinStringSliceDuplicateValue(strSlice []string) bool {
	if len(strSlice) == 0 {
		return false
	}
	_map := make(map[string]struct{}, len(strSlice))
	for _, v := range strSlice {
		if _, ok := _map[v]; ok {
			return true
		}
		_map[v] = struct{}{}
	}
	return false
}

// BuiltinStringSliceRemoveElement 从切片中移除遇到第一个指定元素
func BuiltinStringSliceRemoveElement(source []string, element string) ([]string, bool) {
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
