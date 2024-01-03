/**
@Author: guantingwei
@Email: guantingwei@sixents.com
@Date: 2024/1/3 13:40
@Description:
*/

package maps

// SetDefault 泛型函数, 设置 类似于 Python 的 setdefault 函数
// 它用于操作 map, 为指定的键提供默认值，如果键不存在于映射中。如果键已存在, 则函数返回与键关联的现有值。
// 如果键不存在，则将指定的默认值插入映射，并返回默认值。
// 该函数有三个参数：
//   - m: 要操作的映射。它应该是一个键类型为 K，值类型为 V 的映射。
//   - key: 要设置或获取默认值的键。
//   - defaultValue: 如果键在映射中不存在，则要设置或返回的默认值。
// 示例用法：
//   myMap := make(map[string]int)
//   result := SetDefault(myMap, "a", 42)
//   // 如果 "a" 在 myMap 中存在，result 将是与 "a" 关联的值。
//   // 如果 "a" 不存在，它将以值 42 插入，并且 result 将是 42。
func SetDefault[M ~map[K]V, K comparable, V any](m M, key K, defaultValue V) V {
	if value, exists := m[key]; exists {
		return value
	}
	m[key] = defaultValue
	return defaultValue
}
