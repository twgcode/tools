/**
@Author: wei-g
@Date:   2020/12/17 3:25 下午
@Description:
*/

package simplesql

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNullInt64_MarshalJSON(t *testing.T) {
	type test struct {
		input NullInt64
		want  []byte
	}
	tests := map[string]test{
		"null":  {NullInt64{}, []byte(`null`)},
		"base":  {*NewNullInt64(10, true), []byte(`10`)},
		"false": {*NewNullInt64(10, false), []byte(`null`)},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tc.input)
			if err != nil {
				t.Errorf("input: %#v, 出现异常错误: %s\n", tc.input, err)
				return
			}
			if string(got) != string(tc.want) {
				t.Errorf("excepted: %s, got: %s", tc.want, got)
			}
		})
	}
}

func TestNullInt64_UnmarshalJSON(t *testing.T) {
	type test struct {
		want  NullInt64
		input []byte
	}
	tests := map[string]test{
		"null": {NullInt64{}, []byte(`null`)},
		"base": {*NewNullInt64(10, true), []byte(`10`)},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var got NullInt64
			err := json.Unmarshal(tc.input, &got)
			if err != nil {
				t.Errorf("input: %#v, 出现异常错误: %s\n", tc.input, err)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s excepted:%v, got:%v", name, tc.want, got) // 将测试用例的name格式化输出
				return
			}
		})
	}
}
