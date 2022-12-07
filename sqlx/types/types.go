/**
@Author: wei-g
@Email: guantingwei@sixents.com
@Date: 2022/12/7 17:13
@Description:
*/

package types

import (
	"encoding/json"
	"github.com/jmoiron/sqlx/types"
)

// NewEmptySliceNullJSONText 参加一个空的 切片的  types.NullJSONText
func NewEmptySliceNullJSONText() (result types.NullJSONText) {
	_ = result.UnmarshalJSON([]byte{91, 93}) // 91 93 []
	result.Valid = true
	return
}

// NewSliceNullJSONText 创建 一个 切片的 NullJSONText, 不关心 source 中元素的类型
func NewSliceNullJSONText(source interface{}, empty bool) (result types.NullJSONText, err error) {
	var data []byte
	if source == nil && empty {
		result = NewEmptySliceNullJSONText()
		return
	}
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}

// NewInt32SliceNullJSONText 创建 一个 int32 切片的 NullJSONText
func NewInt32SliceNullJSONText(source []int32, empty bool) (result types.NullJSONText, err error) {
	var data []byte
	if len(source) == 0 && empty {
		result = NewEmptySliceNullJSONText()
		return
	}
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}

// NewInt64SliceNullJSONText 创建 一个 int64 切片的 NullJSONText
func NewInt64SliceNullJSONText(source []int64, empty bool) (result types.NullJSONText, err error) {
	var data []byte
	if len(source) == 0 && empty {
		result = NewEmptySliceNullJSONText()
		return
	}
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}

// NewStringSliceNullJSONText 创建 一个 string 切片的 NullJSONText
func NewStringSliceNullJSONText(source []string, empty bool) (result types.NullJSONText, err error) {
	var data []byte
	if len(source) == 0 && empty {
		result = NewEmptySliceNullJSONText()
		return
	}
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}
