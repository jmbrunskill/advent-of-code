package main

import (
	"fmt"
	"testing"
)

func TestMult2020_3(t *testing.T) {
	tt := []struct {
		input    []int
		expected int
	}{
		{[]int{979, 366, 675}, 241861950},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d vs %d", tc.input, tc.expected), func(t *testing.T) {
			result := mult2020_3(tc.input)

			if tc.expected != result {
				t.Fatalf("expected %d; got %d", tc.expected, result)
			}
		})
	}
}
