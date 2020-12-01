package main

import (
	"fmt"
	"testing"
)

func TestFind2020(t *testing.T) {
	tt := []struct {
		input    []int
		expected int
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, 514579},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d vs %d", tc.input, tc.expected), func(t *testing.T) {
			result := mult2020(tc.input)

			if tc.expected != result {
				t.Fatalf("expected %d; got %d", tc.expected, result)
			}
		})
	}
}
