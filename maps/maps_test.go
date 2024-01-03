/**
@Author: guantingwei
@Email: guantingwei@sixents.com
@Date: 2024/1/3 14:50
@Description:
*/

package maps

import "testing"

func TestSetDefault(t *testing.T) {
	tests := map[string]struct {
		initialMap     map[string]int
		key            string
		defaultValue   int
		expectedResult int
	}{
		"KeyNotExists":   {make(map[string]int), "a", 42, 42},
		"KeyExists":      {map[string]int{"a": 10}, "a", 42, 10},
		"OtherKeysExist": {map[string]int{"b": 99}, "a", 42, 42},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := SetDefault(tc.initialMap, tc.key, tc.defaultValue)
			if result != tc.expectedResult {
				t.Errorf("For test case '%s', expected result to be %d, but got %d", name, tc.expectedResult, result)
			}
		})
	}
}
