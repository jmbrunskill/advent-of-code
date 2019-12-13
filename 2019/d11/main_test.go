package main

import (
	"fmt"
	"testing"
)

func TestPaintPanelCount(t *testing.T) {
	tt := []struct {
		outputs  []int
		expected int
	}{{[]int{1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0}, 6}}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Expected %v", tc.expected), func(t *testing.T) {
			result := paintPanelCountTest(tc.outputs)
			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}
