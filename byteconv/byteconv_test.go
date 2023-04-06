/**
@Author: wei-G
@Email: wei_g_it@163.com
@Date: 2023/4/6 10:19
@Description:
*/

package byteconv

import (
	"testing"
)

func TestGBToByte(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  int64
	}{
		{"1", 1, 1000000000},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := GBToByte(tt.input)
			if got != tt.want {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}
