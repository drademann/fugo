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

func Reduce[E any](slice []E, reduceFn func(E, E) E) E {
	if len(slice) == 0 {
		panic("empty slice")
	}
	result := slice[0]
	for i := 1; i < len(slice); i++ {
		result = reduceFn(result, slice[i])
	}
	return result
}

func ReduceInit[E any](initValue E, slice []E, reduceFn func(E, E) E) E {
	return Reduce(append([]E{initValue}, slice...), reduceFn)
}
