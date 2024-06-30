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

package fp

// Filter takes a slice and a filter function and returns a new slice containing only the elements that pass the filter.
// The function preserves the order of the elements in the original slice and returns an empty slice if the input slice is nil or empty.
func Filter[E any](slice []E, filterFn func(E) bool) []E {
	filtered := make([]E, 0)
	for _, item := range slice {
		if filterFn(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
