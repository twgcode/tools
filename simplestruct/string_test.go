/**
@Author: wei-g
@Date:   2021/8/24 5:24 下午
@Description:
*/

package simplestruct

import (
	"reflect"
	"testing"
)

func TestStringRemoveDuplicateValues(t *testing.T) {
	type test struct {
		input []string
		want  []string
	}
	tests := map[string]test{
		"nil":    {nil, nil},
		"empty":  {[]string{}, []string{}},
		"simple": {[]string{"a", "b", "c", "a"}, []string{"a", "b", "c"}},
		"zh":     {[]string{"a", "b", "中文", "中文", "繁体"}, []string{"a", "b", "中文", "繁体"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := BuiltinStringSliceRemoveDuplicateValues(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(" excepted:%#v, got:%#v", tc.want, got)
				return
			}
		})
	}
}

func TestBuiltinStringSliceRemoveElement(t *testing.T) {
	type test struct {
		input   []string
		element string
		want    []string
		exist   bool
	}
	tests := map[string]test{
		"nil":              {nil, "", nil, false},
		"empty":            {[]string{}, "", []string{}, false},
		"simple":           {[]string{"a", "b", "c", "a"}, "a", []string{"b", "c", "a"}, true},
		"not_exist_simple": {[]string{"a", "b", "c", "a"}, "f", []string{"a", "b", "c", "a"}, false},
		"zh":               {[]string{"中午", "是个", "好的", "a"}, "中午", []string{"是个", "好的", "a"}, true},
		"not_exist_zh":     {[]string{"中午", "是个", "好的", "a"}, "nb", []string{"中午", "是个", "好的", "a"}, false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, ok := BuiltinStringSliceRemoveElement(tc.input, tc.element)
			if ok != tc.exist {
				t.Errorf("excepted slice:%#v, got slice:%#v, excepted: %t", tc.want, got, tc.exist)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("获取到的值错误 excepted slice:%#v, got slice:%#v, excepted: %t", tc.want, got, tc.exist)
				return
			}
		})
	}
}
