/**
@Author: wei-g
@Date:   2020/11/13 10:32 上午
@Description:
*/

package simplestruct

import (
	"encoding/json"
	"reflect"
	"testing"
)

type TestFoo struct {
	Bars []Int64 `json:"bars"`
	Name string  `json:"name"`
	ID   Int64   `json:"id"`
}

// TestInt64Str_MarshalJSON 测试 Int64 在进行json 序列化时 是否符合预期
func TestInt64Str_MarshalJSON(t *testing.T) {
	type test struct { // 定义test结构体
		want  string
		input interface{}
	}
	tests := map[string]test{
		"base": {want: `"12"`,
			input: Int64(12),
		},
		"negative": {want: `"-12"`,
			input: Int64(-12),
		},
		"slice": {want: `["1","2","3"]`,
			input: []Int64{1, 2, 3},
		},
		"struct": {want: `{"bars":["1","2","3","4"],"name":"one"}`,
			input: struct {
				Bars []Int64 `json:"bars"`
				Name string  `json:"name"`
			}{
				Bars: []Int64{1, 2, 3, 4},
				Name: "one"},
		},
		"struct_simple": {
			want: `{"bars":"1","name":"one"}`,
			input: struct {
				Bars Int64  `json:"bars"`
				Name string `json:"name"`
			}{
				Bars: Int64(1),
				Name: "one"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tc.input)
			if err != nil {
				t.Errorf("input: %s, 出现异常错误: %s\n", tc.input, err)
				return
			}
			if string(got) != tc.want {
				t.Errorf("excepted: %s, got: %s", tc.want, got)
			}
		})
	}
}

// TestInt64Str_UnmarshalJSON 测试 Int64 在进行json 反序列化时 是否符合预期
func TestInt64Str_UnmarshalJSON(t *testing.T) {
	type test struct { // 定义test结构体
		want  Int64
		input string
	}

	tests := map[string]test{
		"base":            {want: Int64(12), input: `12`},
		"symbol":          {want: Int64(12), input: `"12"`},
		"negative":        {want: Int64(-12), input: `-12`},
		"symbol_negative": {want: Int64(-12), input: `"-12"`},
		"big":             {want: Int64(9223372036854775807), input: "9223372036854775807"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var i Int64
			err := json.Unmarshal([]byte(tc.input), &i)
			if err != nil {
				t.Errorf("input: %s, 出现异常错误: %s\n", tc.input, err)
				return
			}
			if i != tc.want {
				t.Errorf("excepted: %d, got: %d", tc.want, i)
			}
		})
	}

}

// TestInt64_UnmarshalJSONTestFoo 测试 Int64 在结构体中序列化 表现
func TestInt64_UnmarshalJSONTestFoo(t *testing.T) {
	type test struct {
		input string
		want  TestFoo
	}
	tests := map[string]test{
		"nil_slice": {
			input: `{"bars":null,"name":"one","id":1}`,
			want: TestFoo{
				Bars: nil,
				Name: "one",
				ID:   1,
			},
		},
		"empty_slice": {
			input: `{"bars":[],"name":"two","id":"2"}`,
			want: TestFoo{
				Bars: []Int64{},
				Name: "two",
				ID:   2,
			},
		},
		"slice": {
			input: `{"bars":["1","2","3", "4",5],"name":"three","id":"-22000"}`,
			want: TestFoo{
				Bars: []Int64{1, 2, 3, 4, 5},
				Name: "three",
				ID:   -22000,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var got TestFoo
			err := json.Unmarshal([]byte(tc.input), &got)
			if err != nil {
				t.Errorf("input: %s, 出现异常错误: %s\n", tc.input, err)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(" excepted:%#v, got:%#v", tc.want, got)
			}
			//t.Logf("got: %#v\n", got)
		})
	}

}

func TestToBuiltinInt64Slice(t *testing.T) {
	type test struct {
		input []Int64
		want  []int64
	}
	tests := map[string]test{
		"nil":   {nil, nil},
		"empty": {[]Int64{}, []int64{}},
		"short": {[]Int64{1, 2, 3, 4, 5}, []int64{1, 2, 3, 4, 5}},
		"long":  {make([]Int64, 1024, 2048), make([]int64, 1024, 2048)},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ToBuiltinInt64Slice(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(" excepted:%#v, got:%#v", tc.want, got)
				return
			}
			if cap(got) != cap(tc.want) {
				t.Errorf("cap unequal, excepted: %d, got: %d", cap(tc.want), cap(got))
				return
			}
		})
	}
}

func TestToInt64Slice(t *testing.T) {
	type test struct {
		input []int64
		want  []Int64
	}
	tests := map[string]test{
		"nil":   {nil, nil},
		"empty": {[]int64{}, []Int64{}},
		"short": {[]int64{1, 2, 3, 4, 5}, []Int64{1, 2, 3, 4, 5}},
		"long":  {make([]int64, 1024, 2048), make([]Int64, 1024, 2048)},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ToInt64Slice(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(" excepted:%#v, got:%#v", tc.want, got)
				return
			}
			if cap(got) != cap(tc.want) {
				t.Errorf("cap unequal, excepted: %d, got: %d", cap(tc.want), cap(got))
				return
			}
		})
	}
}

func TestAppendIntSlice(t *testing.T) {
	type test struct { // 定义test结构体
		want   []int64
		input1 []int64
		input2 []int64
	}
	tests := map[string]test{
		"first_nil":   {want: []int64{1, 2, 3}, input1: nil, input2: []int64{1, 2, 3}},
		"second_nil":  {want: []int64{1, 2, 3}, input1: []int64{1, 2, 3}, input2: nil},
		"all_nil":     {want: nil, input1: nil, input2: nil},
		"empty_nil":   {want: []int64{}, input1: []int64{}, input2: nil},
		"all_empty":   {want: []int64{}, input1: []int64{}, input2: []int64{}},
		"short_empty": {want: []int64{1, 2, 3}, input1: []int64{1, 2, 3}, input2: []int64{}},
		"all_short":   {want: []int64{1, 2, 3, 1, 2, 3, 4}, input1: []int64{1, 2, 3}, input2: []int64{1, 2, 3, 4}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := AppendBuiltinInt64Slice(tc.input1, tc.input2)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(" excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
