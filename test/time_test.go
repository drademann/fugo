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

package test_test

import (
	"fugo/test"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	date := test.Date("30.06.2024 10:20")

	expectedDate := time.Date(2024, time.June, 30, 10, 20, 0, 0, time.Local)
	if date != expectedDate {
		t.Errorf("expected %v, got %v", expectedDate, date)
	}
}

func TestTime(t *testing.T) {
	tm := test.Time("10:20")

	expectedTime := time.Date(2024, time.June, 21, 10, 20, 0, 0, time.Local)
	if tm != expectedTime {
		t.Errorf("expected %v, got %v", expectedTime, tm)
	}
}
