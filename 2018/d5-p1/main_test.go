package main

import (
	"fmt"
	"testing"
)

func TestProcessInput(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
		{"aA", ""},
		{"abBA", ""},
		{"aabAAB", "aabAAB"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf(tc.input, tc.input, tc.expected), func(t *testing.T) {

			str := react(tc.input)

			if tc.expected != str {
				t.Fatalf("expected %s; got %s", tc.expected, str)
			}
		})
	}
}
