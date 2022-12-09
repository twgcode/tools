/**
@Author: wei-G
@Email: 17600113577@163.com
@Date: 2022/12/9 14:41
@Description:
*/

package simplesql

import "database/sql"

// NewNullInt32FromInt32Ptr 根据  *int32 创建 sql.NullInt32
func NewNullInt32FromInt32Ptr(source *int32) (result sql.NullInt32) {
	if source == nil {
		result.Valid = false
		return
	}
	result.Valid = true
	result.Int32 = *source
	return
}
