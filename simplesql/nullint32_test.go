/**
@Author: wei-G
@Email: 17600113577@163.com
@Date: 2022/12/9 14:42
@Description:
*/

package simplesql

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestNewNullInt32FromInt32Ptr(t *testing.T) {
	type test struct {
		input *int32
		want  sql.NullInt32
	}
	var normal int32 = 10
	tests := map[string]test{
		"null":   {nil, sql.NullInt32{Int32: 0, Valid: false}},
		"normal": {&normal, sql.NullInt32{Int32: 10, Valid: true}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewNullInt32FromInt32Ptr(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%v, got:%v", tc.want, got) // 将测试用例的name格式化输出
				return
			}
		})
	}
}
