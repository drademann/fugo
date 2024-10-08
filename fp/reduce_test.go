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
	"testing"
)

func TestReduce(t *testing.T) {
	aSlice := []int{1, 2, 3, 4, 5}

	sum := fp.Reduce(aSlice, func(acc int, n int) int { return acc + n })

	if sum != 15 {
		t.Errorf("expected 15, got %d", sum)
	}
}

func TestReduceInit(t *testing.T) {
	aSlice := []int{1, 2, 3, 4, 5}

	sum := fp.ReduceInit(10, aSlice, func(acc int, n int) int { return acc + n })

	if sum != 25 {
		t.Errorf("expected 25, got %d", sum)
	}
}

func TestReduceInit_givenEmptySlice_shouldReturnInitValue(t *testing.T) {
	sum := fp.ReduceInit(7, []int{}, func(acc int, n int) int { return acc + n })

	if sum != 7 {
		t.Errorf("expected 7, got %d", sum)
	}
}
