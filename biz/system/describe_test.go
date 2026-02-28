package system

import (
	"maps"
	"testing"
)

func TestDescribeByRuleSet(t *testing.T) {
	dict := map[string]string{"a": "1"}
	it := maps.Keys(dict)
	describeX := DescribeByRuleSet(func(key string) string {
		if value, ok := dict[key]; ok {
			return value
		} else {
			return "N/A"
		}
	})
	for v := range it {
		t.Run(v, func(tc *testing.T) {
			tc.Log(describeX(v))
		})

	}

}
