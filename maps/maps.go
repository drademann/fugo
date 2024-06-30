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

package maps

import (
	"cmp"
	"sort"
)

// Keys returns a slice of keys extracted from the input map.
func Keys[K comparable, E any](aMap map[K]E) []K {
	var keys []K
	for k := range aMap {
		keys = append(keys, k)
	}
	return keys
}

// SortedKeys returns a sorted slice of keys extracted from the input map.
func SortedKeys[K cmp.Ordered, E any](aMap map[K]E) []K {
	keys := Keys(aMap)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}
