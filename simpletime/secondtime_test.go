/**
@Author: wei-g
@Date:   2020/11/25 9:31 下午
@Description:
*/

package simpletime

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

type Post struct {
	CreateTime SecondTime `json:"create_time"`
	Name       string     `json:"name"`
}

func TestTime_IsSet(t *testing.T) {
	type test struct {
		input SecondTime
		want  bool
	}
	tests := map[string]test{
		"nil_time": {SecondTime{time.Time{}}, false},
		"now":      {SecondTime{time.Now()}, true},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.IsSet()
			if got != tc.want {
				t.Errorf(" excepted:%t, got:%t", tc.want, got)
			}
		})
	}
}
func TestTime_MarshalJSON(t *testing.T) {
	type test struct {
		input SecondTime
		want  string
	}
	// PRC 中国时区
	location, _ := time.LoadLocation("PRC")

	tests := map[string]test{
		"nil_time": {SecondTime{time.Time{}}, "null"},
		"CST":      {SecondTime{time.Date(2020, 11, 25, 23, 24, 53, 0, location)}, `"2020-11-25 23:24:53 CST"`},
		"UTC":      {SecondTime{time.Date(2020, 11, 25, 23, 24, 53, 0, time.UTC)}, `"2020-11-25 23:24:53 UTC"`},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(&tc.input)
			if err != nil {
				t.Errorf("input %v, Marshal 出现异常错误: %s", tc.input, err)
				return
			}
			if string(got) != tc.want {
				t.Errorf(" excepted:%s, got:%s", tc.want, got)
				return
			}
		})
	}
}

func TestTime_MarshalJSONStruct(t *testing.T) {
	type test struct {
		input Post
		want  string
	}

	tests := map[string]test{
		"nil_time": {Post{
			CreateTime: SecondTime{},
			Name:       "go",
		}, `{"create_time":null,"name":"go"}`},
		"utc": {
			Post{SecondTime{time.Date(2020, 11, 25, 23, 24, 53, 0, time.UTC)}, "utc"},
			`{"create_time":"2020-11-25 23:24:53 UTC","name":"utc"}`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(&tc.input)
			if err != nil {
				t.Errorf("input %v, Marshal 出现异常错误: %s", tc.input, err)
				return
			}
			if string(got) != tc.want {
				t.Errorf(" excepted:%s, got:%s", tc.want, got)
				return
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	type test struct {
		input string
		want  SecondTime
	}
	tests := map[string]test{
		"nil_time": {"null", SecondTime{}},
		"uct":      {`"2020-11-25 23:24:53 UTC"`, SecondTime{time.Date(2020, 11, 25, 23, 24, 53, 0, time.UTC)}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var got SecondTime
			err := json.Unmarshal([]byte(tc.input), &got)
			if err != nil {
				t.Errorf("input %v, Marshal 出现异常错误: %s", tc.input, err)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
			}
		})
	}
}

func TestTime_UnmarshalJSON2Struct(t *testing.T) {
	type test struct {
		input string
		want  Post
	}
	tests := map[string]test{
		"nil_time": {`{"create_time":null,"name":"go"}`,
			Post{CreateTime: SecondTime{}, Name: "go"}},
		"utc": {`{"create_time":"2020-11-25 23:24:53 UTC","name":"utc"}`,
			Post{SecondTime{time.Date(2020, 11, 25, 23, 24, 53, 0, time.UTC)}, "utc"},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var got Post
			err := json.Unmarshal([]byte(tc.input), &got)
			if err != nil {
				t.Errorf("input %v, Marshal 出现异常错误: %s", tc.input, err)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
			}
		})
	}
}
