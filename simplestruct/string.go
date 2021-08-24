/**
@Author: wei-g
@Date:   2021/8/24 5:20 下午
@Description:
*/

package simplestruct

// StringContains 判断 item是发在指定的切片中
func StringContains(slice []string, val string) (ok bool) {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// StringRemoveDuplicateValues 对切片进行去除
func StringRemoveDuplicateValues(strSlice []string) []string {
	if len(strSlice) == 0 {
		return strSlice
	}
	list := make([]string, 0, cap(strSlice))
	for _, item := range strSlice {
		if !StringContains(list, item) {
			list = append(list, item)
		}
	}
	return list
}
