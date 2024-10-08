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

package assert

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"
)

// ErrorString is a function that compares the error string representation with the expected string.
func ErrorString(t *testing.T, err error, expected string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected an error %q but got none", expected)
	}
	if err.Error() != expected {
		t.Errorf("expected error %q but got %q", expected, err.Error())
	}
}

// ErrorIs is a function to check if the given error matches the expected error using errors.Is().
func ErrorIs(t *testing.T, err, expected error) {
	t.Helper()
	if !errors.Is(err, expected) {
		t.Errorf("expected error %v but got %v", expected, err)
	}
}

// NoError is a function to check if the given error is nil. If the error is not nil, it raises an error.
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// Output is a function to compare a given output with an expected output and raise an error if they don't match.
// It removes any leading newline character from the expected output and removes any tab characters.
func Output(t *testing.T, out *bytes.Buffer, expected string) {
	t.Helper()
	if strings.HasPrefix(expected, "\n") {
		expected = expected[1:]
	}
	expected = strings.ReplaceAll(expected, "\t", "")
	if out.String() != expected {
		t.Errorf("expected output \n%q but got \n%q", expected, out.String())
	}
}

// Duration is a function to compare a given duration with an expected one and raise an error if they don't match.
func Duration(t *testing.T, d, expected time.Duration) {
	t.Helper()
	if d != expected {
		t.Errorf("expected duration of %v but got %v", expected, d)
	}
}

// DurationString is a function to compare a given duration with an expected duration defined as a string.
// The string representation follows the standard format, for example, “15m”.
// The function adds zero values automatically at the end of the expected duration string:
// Given “15m” uses “15m0s” for comparison.
func DurationString(t *testing.T, d time.Duration, expected string) {
	t.Helper()
	if strings.HasSuffix(expected, "h") {
		expected = expected + "0m"
	}
	if strings.HasSuffix(expected, "m") {
		expected = expected + "0s"
	}
	if d.String() != expected {
		t.Errorf("expected duration string of %q but got %q", expected, d.String())
	}
}

// ContainsInAnyOrder is a function to compare a given actual slice with an expected slice,
// checking if the actual slice contains all the elements in the expected slice in any order.
func ContainsInAnyOrder[E comparable](t *testing.T, actual []E, expected []E) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Errorf("expected %d elements but got %d", len(expected), len(actual))
		return
	}
	for _, expectedValue := range expected {
		if !slices.Contains(actual, expectedValue) {
			t.Errorf("expected %v to contain all of %v", actual, expected)
			return
		}
	}
}

func True(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Error("expected true but got false")
	}
}

func False(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		t.Error("expected false but got true")
	}
}

// Equal compares two comparable values
func Equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %s but got %s", toString(expected), toString(actual))
	}
}

func toString[T comparable](v T) string {
	switch val := any(v).(type) {
	case string:
		return fmt.Sprintf("%q", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// EqualSlice compares the elements of two slices for equality.
func EqualSlice[E comparable](t *testing.T, actual []E, expected []E) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Errorf("expected slice %v but got %v with different length %d", expected, actual, len(actual))
	}
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("expected slice %v but got %v", expected, actual)
			return
		}
	}
}
