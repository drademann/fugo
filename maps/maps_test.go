/*
 *
 * Copyright 2024-2024 The Fugo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package maps_test

import (
	"github.com/drademann/fugo/maps"
	"github.com/drademann/fugo/test/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	aMap := map[string]int{"a": 3, "c": 2, "b": 1}

	keys := maps.Keys(aMap)

	expectedKeys := []string{"a", "b", "c"}
	assert.ContainsInAnyOrder(t, keys, expectedKeys)
}

func TestSortedKeys(t *testing.T) {
	aMap := map[string]int{"a": 3, "c": 2, "b": 1}

	keys := maps.SortedKeys(aMap)
	assert.EqualSlice(t, keys, []string{"a", "b", "c"})
}
