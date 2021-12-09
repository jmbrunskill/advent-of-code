package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	tt := []struct {
		input    []int
		days     int
		expected int
	}{
		{[]int{3, 4, 3, 1, 2}, 18, 26},
		{[]int{3, 4, 3, 1, 2}, 80, 5934},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d vs %d", tc.input, tc.expected), func(t *testing.T) {
			result := simFish(tc.input, tc.days)

			if tc.expected != result {
				t.Fatalf("expected %d; got %d", tc.expected, result)
			}
		})
	}
}
