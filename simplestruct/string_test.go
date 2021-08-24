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
			got := StringRemoveDuplicateValues(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(" excepted:%#v, got:%#v", tc.want, got)
				return
			}
		})
	}
}
