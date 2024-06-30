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

package fp_test

import (
	"github.com/drademann/fugo/fp"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	aSlice := []int{1, 2, 3, 4, 5}

	filtered := fp.Filter(aSlice, func(x int) bool { return x%2 != 0 })

	expected := []int{1, 3, 5}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("expected %v but got %v", expected, filtered)
	}
}
