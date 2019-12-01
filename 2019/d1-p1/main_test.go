package main

import (
	"fmt"
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	tt := []struct {
		input    int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d vs %d", tc.input, tc.expected), func(t *testing.T) {
			result := calculateFuel(tc.input)

			if tc.expected != result {
				t.Fatalf("expected %d; got %d", tc.expected, result)
			}
		})
	}
}
