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
		{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", "138"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s vs %s", tc.input, tc.expected), func(t *testing.T) {

			str := processInput(tc.input)

			if tc.expected != str {
				t.Fatalf("expected %s; got %s", tc.expected, str)
			}
		})
	}
}
