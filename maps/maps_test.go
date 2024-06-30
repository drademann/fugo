package maps_test

import (
	"fugo/maps"
	"fugo/test/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	aMap := map[string]int{"a": 3, "c": 2, "b": 1}

	keys := maps.Keys(aMap)

	expectedKeys := []string{"a", "b", "c"}
	assert.ContainsInAnyOrder(t, keys, expectedKeys)
}
