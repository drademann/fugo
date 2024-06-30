/*
 *
 * Copyright 2024-2024 The Haora Authors
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

package assert_test

import (
	"bytes"
	"fugo/test/assert"
	"testing"
	"time"
)

func TestAssertOutput(t *testing.T) {
	out := bytes.NewBufferString("Hello World\n")

	assert.Output(t, out,
		`
		Hello World
		`)
}

func ExampleOutput() {
	t := &testing.T{} // example testing

	assert.Output(t, bytes.NewBufferString("Hello World\n"),
		`
		Hello World
		`)
}

func TestAssertDuration(t *testing.T) {
	assert.Duration(t, 15*time.Minute, 15*time.Minute)
}

func TestAssertDurationString(t *testing.T) {
	assert.DurationString(t, 1*time.Hour, "1h")
	assert.DurationString(t, 15*time.Minute, "15m")
	assert.DurationString(t, 30*time.Second, "30s")
}

func ExampleDurationString() {
	t := &testing.T{} // example testing

	assert.DurationString(t, 1*time.Hour, "1h")
	assert.DurationString(t, 1*time.Hour, "1h0m0s")
	assert.DurationString(t, 15*time.Minute, "15m")
}
