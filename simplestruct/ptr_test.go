/**
@Author: wei-G
@Email: 17600113577@163.com
@Date: 2022/12/12 10:12
@Description:
*/

package simplestruct

import (
	"reflect"
	"testing"
)

func TestPtrStringValue(t *testing.T) {
	type test struct { // 定义test结构体
		want  string
		input *string
	}
	emptyStr := ""
	enStr := "abc"
	hanziStr := "中wen测试"
	tests := map[string]test{

		"nil":   {"", nil},
		"empty": {emptyStr, &emptyStr},
		"en":    {enStr, &enStr},
		"hanzi": {hanziStr, &hanziStr},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PtrStringValue(tc.input)
			if got != tc.want {
				t.Errorf("excepted: %q, got: %q", tc.want, got)
			}
		})
	}
}

func TestPtrSliceStringValue(t *testing.T) {
	type test struct { // 定义test结构体
		want  []string
		input []*string
		flag  bool
	}
	a := "a"
	b := "b"
	tests := map[string]test{
		"nil":                {[]string{}, nil, true},
		"nil_false":          {[]string{}, nil, false},
		"empty":              {[]string{}, []*string{}, true},
		"empty_false":        {[]string{}, []*string{}, false},
		"contains nil":       {[]string{"a", "b"}, []*string{&a, nil, &b}, true},
		"contains nil false": {[]string{"a", "", "b"}, []*string{&a, nil, &b}, true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := PtrSliceStringValue(tc.input, tc.flag)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("excepted: %s, got: %s", tc.want, got)
			}

		})
	}

}
