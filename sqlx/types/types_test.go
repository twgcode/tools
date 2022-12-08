/**
@Author: wei-g
@Email: guantingwei@sixents.com
@Date: 2022/12/7 17:22
@Description:
*/

package types

import (
	"testing"
)

func TestNewEmptySliceNullJSONText(t *testing.T) {
	var (
		excepted = "[]"
	)
	data := NewEmptySliceNullJSONText()
	got := data.String()
	if got != excepted {
		t.Errorf(" excepted:%s, got: %s\n", excepted, got)
	}
}

func TestNewStringSliceNullJSONText(t *testing.T) {
	type test struct {
		input []string
		want  string
	}
	tests := map[string]test{
		"nil":     {nil, "[]"},
		"english": {[]string{"a", "b", "c"}, `["a","b","c"]`},
		"hanzi":   {[]string{"您好", "多好", "中文", "测试"}, `["您好","多好","中文","测试"]`},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			nullJSONText, err := NewStringSliceNullJSONText(tc.input, true)
			if err != nil {
				t.Errorf("序列化失败: %s\n", err)
				return
			}
			marshalJSON, err := nullJSONText.MarshalJSON()
			if err != nil {
				t.Errorf("获取序列化后json字符串失败: %s\n", err)
				return
			}
			marshalJSONStr := string(marshalJSON)
			if marshalJSONStr != tc.want {
				t.Errorf(" excepted:%s, got:%s", tc.want, marshalJSONStr)
				return
			}
		})
	}
}

type User struct {
	Name string `json:"name"`
}

func TestNewSliceNullJSONText(t *testing.T) {
	type test struct {
		input interface{}
		want  string
	}
	var strSlicePtrNil []string
	var strSlicePtrEmpty = make([]string, 0, 0)
	var structSlice = []User{{"x"}}

	tests := map[string]test{
		"nil":              {input: nil, want: "[]"},
		"strSlicePtrNil":   {input: strSlicePtrNil, want: "[]"},
		"strSlice":         {input: []string{"A", "B"}, want: `["A","B"]`},
		"strSlicePtr":      {input: &[]string{"A", "B"}, want: `["A","B"]`},
		"strSlicePtrEmpty": {input: &strSlicePtrEmpty, want: `[]`},
		"structSlice":      {input: structSlice, want: `[{"name":"x"}]`},
		"structSlicePtr":   {input: &structSlice, want: `[{"name":"x"}]`},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			nullJSONText, err := NewSliceNullJSONText(tc.input, true)
			if err != nil {
				t.Errorf("序列化失败: %s\n", err)
				return
			}
			marshalJSON, err := nullJSONText.MarshalJSON()
			if err != nil {
				t.Errorf("获取序列化后json字符串失败: %s\n", err)
				return
			}
			marshalJSONStr := string(marshalJSON)
			if marshalJSONStr != tc.want {
				t.Errorf(" excepted:%s, got:%s", tc.want, marshalJSONStr)
				return
			}
		})
	}
}

func TestNewSliceJSONText(t *testing.T) {
	type test struct {
		input interface{}
		want  string
	}
	var strSlicePtrNil []string
	var strSlicePtrEmpty = make([]string, 0, 0)
	var structSlice = []User{{"x"}}

	tests := map[string]test{
		"nil":              {input: nil, want: "[]"},
		"strSlicePtrNil":   {input: strSlicePtrNil, want: "[]"},
		"strSlice":         {input: []string{"A", "B"}, want: `["A","B"]`},
		"strSlicePtr":      {input: &[]string{"A", "B"}, want: `["A","B"]`},
		"strSlicePtrEmpty": {input: &strSlicePtrEmpty, want: `[]`},
		"structSlice":      {input: structSlice, want: `[{"name":"x"}]`},
		"structSlicePtr":   {input: &structSlice, want: `[{"name":"x"}]`},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			nullJSONText, err := NewSliceJSONText(tc.input, true)
			if err != nil {
				t.Errorf("序列化失败: %s\n", err)
				return
			}
			marshalJSON, err := nullJSONText.MarshalJSON()
			if err != nil {
				t.Errorf("获取序列化后json字符串失败: %s\n", err)
				return
			}
			marshalJSONStr := string(marshalJSON)
			if marshalJSONStr != tc.want {
				t.Errorf(" excepted:%s, got:%s", tc.want, marshalJSONStr)
				return
			}
		})
	}
}
