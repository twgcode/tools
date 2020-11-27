/**
@Author: wei-g
@Date:   2020/11/26 7:18 下午
@Description:
*/

package simplesql

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestSecondNullTime_MarshalJSON(t *testing.T) {
	type test struct {
		input SecondNullTime
		want  string
	}
	// PRC 中国时区
	location, _ := time.LoadLocation("PRC")

	tests := map[string]test{
		"nil_time": {ToSecondNullTime(time.Time{}), "null"},
		"CST":      {ToSecondNullTime(time.Date(2020, 11, 25, 23, 24, 53, 0, location)), `"2020-11-25 23:24:53 CST"`},
		"UTC":      {ToSecondNullTime(time.Date(2020, 11, 25, 23, 24, 53, 0, time.UTC)), `"2020-11-25 23:24:53 UTC"`},
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

func TestSecondNullTime_UnmarshalJSON(t *testing.T) {
	type test struct {
		input string
		want  SecondNullTime
	}
	tests := map[string]test{
		"nil_time": {"null",
			SecondNullTime{}},
		"uct": {`"2020-11-25 23:24:53 UTC"`,
			ToSecondNullTime(time.Date(2020, 11, 25, 23, 24, 53, 0, time.UTC)),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var got SecondNullTime
			err := json.Unmarshal([]byte(tc.input), &got)
			if err != nil {
				t.Errorf("input %v, Marshal 出现异常错误: %s", tc.input, err)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s excepted:%v, got:%v", name, tc.want, got) // 将测试用例的name格式化输出
				return
			}
			t.Log(got)
		})
	}
}
