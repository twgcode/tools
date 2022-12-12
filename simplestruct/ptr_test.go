/**
@Author: wei-G
@Email: 17600113577@163.com
@Date: 2022/12/12 10:12
@Description:
*/

package simplestruct

import (
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
