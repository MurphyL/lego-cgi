package desc

import (
	"maps"
	"testing"
)

func TestViaMap(t *testing.T) {
	dict := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	desc := ViaMap[int](dict)
	for v := range maps.Keys(dict) {
		t.Run(v, func(tc *testing.T) {
			tc.Log(desc(v))
		})

	}
}
