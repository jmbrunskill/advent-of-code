package main

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {
	tt := []struct {
		input string
		min   int
		max   int
		c     string
		pw    string
	}{
		{"1-3 a: abcde", 1, 3, "a", "abcde"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.input, tc.pw), func(t *testing.T) {
			min, max, c, pw := parseLine(tc.input)

			if tc.min != min {
				t.Fatalf("expected %v; got %v", tc.min, min)
			}
			if tc.max != max {
				t.Fatalf("expected %v; got %v", tc.max, max)
			}
			if tc.c != c {
				t.Fatalf("expected %v; got %v", tc.c, c)
			}
			if tc.pw != pw {
				t.Fatalf("expected %v; got %v", tc.pw, pw)
			}
		})
	}
}
func TestValidatePassword(t *testing.T) {
	tt := []struct {
		input    string
		expected bool
	}{
		{"1-3 a: abcde", true},
		{"1-3 b: cdefg", false},
		{"2-9 c: ccccccccc", false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.input, tc.expected), func(t *testing.T) {
			result := validatePassword(parseLine(tc.input))

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}
