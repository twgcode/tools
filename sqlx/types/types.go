/**
@Author: wei-g
@Email: 17600113577@163.com
@Date: 2022/12/7 17:13
@Description:
*/

package types

import (
	"encoding/json"
	"github.com/jmoiron/sqlx/types"
	"reflect"
)

// NewEmptySliceNullJSONText 参加一个空的 切片的  types.NullJSONText
func NewEmptySliceNullJSONText() (result types.NullJSONText) {
	_ = result.UnmarshalJSON([]byte{91, 93}) // 91 93 []
	result.Valid = true
	return
}

// NewSliceNullJSONText 创建 一个 切片的 types.NullJSONText, 不关心 source 中元素的类型
func NewSliceNullJSONText(source interface{}, empty bool) (result types.NullJSONText, err error) {
	// 如果 source 直接是nil
	if empty && source == nil {
		result = NewEmptySliceNullJSONText()
		return
	}
	sourceType := reflect.TypeOf(source)
	sourceTypeKind := sourceType.Kind()

	// source 为 切片, 比如 var source []string
	if empty && (sourceTypeKind == reflect.Slice || sourceTypeKind == reflect.Array) {
		if reflect.ValueOf(source).Len() == 0 {
			result = NewEmptySliceNullJSONText()
			return
		}
	}
	// source 为指针 但是，指向的是 切片 比如  var source []string 调用时 NewSliceNullJSONText(&source)
	// 或者 var source *[]string
	if sourceTypeKind == reflect.Ptr {
		sourceTypeKind = sourceType.Elem().Kind()
		if empty && (sourceTypeKind == reflect.Slice || sourceTypeKind == reflect.Array) {
			if reflect.ValueOf(source).Elem().Len() == 0 {
				result = NewEmptySliceNullJSONText()
				return
			}
		}
	}

	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}

// NewInt32SliceNullJSONText 创建 一个 int32 切片的 types.NullJSONText
func NewInt32SliceNullJSONText(source []int32, empty bool) (result types.NullJSONText, err error) {
	if len(source) == 0 && empty {
		result = NewEmptySliceNullJSONText()
		return
	}
	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}

// NewInt64SliceNullJSONText 创建 一个 int64 切片的 types.NullJSONText
func NewInt64SliceNullJSONText(source []int64, empty bool) (result types.NullJSONText, err error) {
	if len(source) == 0 && empty {
		result = NewEmptySliceNullJSONText()
		return
	}
	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}

// NewStringSliceNullJSONText 创建 一个 string 切片的 types.NullJSONText
func NewStringSliceNullJSONText(source []string, empty bool) (result types.NullJSONText, err error) {
	if len(source) == 0 && empty {
		result = NewEmptySliceNullJSONText()
		return
	}
	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	result.Valid = true
	return
}

// NewEmptySliceJSONText 参加一个空的 切片的  types.JSONText
func NewEmptySliceJSONText() (result types.JSONText) {
	_ = result.UnmarshalJSON([]byte{91, 93}) // 91 93 []
	return
}

// NewSliceJSONText 创建 一个 切片的 types.NullJSONText, 不关心 source 中元素的类型
func NewSliceJSONText(source interface{}, empty bool) (result types.JSONText, err error) {
	// 如果 source 直接是nil
	if empty && source == nil {
		result = NewEmptySliceJSONText()
		return
	}
	sourceType := reflect.TypeOf(source)
	sourceTypeKind := sourceType.Kind()

	// source 为 切片, 比如 var source []string
	if empty && (sourceTypeKind == reflect.Slice || sourceTypeKind == reflect.Array) {
		if reflect.ValueOf(source).Len() == 0 {
			result = NewEmptySliceJSONText()
			return
		}
	}
	// source 为指针 但是，指向的是 切片 比如  var source []string 调用时 NewSliceNullJSONText(&source)
	// 或者 var source *[]string
	if sourceTypeKind == reflect.Ptr {
		sourceTypeKind = sourceType.Elem().Kind()
		if empty && (sourceTypeKind == reflect.Slice || sourceTypeKind == reflect.Array) {
			if reflect.ValueOf(source).Elem().Len() == 0 {
				result = NewEmptySliceJSONText()
				return
			}
		}
	}

	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	return
}

// NewInt32SliceJSONText 创建 一个 int32 切片的 NullJSONText
func NewInt32SliceJSONText(source []int32, empty bool) (result types.JSONText, err error) {
	if len(source) == 0 && empty {
		result = NewEmptySliceJSONText()
		return
	}
	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}

	return
}

// NewInt64SliceJSONText 创建 一个 int64 切片的 NullJSONText
func NewInt64SliceJSONText(source []int64, empty bool) (result types.JSONText, err error) {
	if len(source) == 0 && empty {
		result = NewEmptySliceJSONText()
		return
	}
	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	return
}

// NewStringSliceJSONText 创建 一个 string 切片的 NullJSONText
func NewStringSliceJSONText(source []string, empty bool) (result types.JSONText, err error) {
	if len(source) == 0 && empty {
		result = NewEmptySliceJSONText()
		return
	}
	var data []byte
	if data, err = json.Marshal(source); err != nil {
		return
	}
	if err = result.UnmarshalJSON(data); err != nil {
		return
	}
	return
}
