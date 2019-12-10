package main

import (
	"fmt"
	"testing"
)

func TestCheckCriteria(t *testing.T) {
	tt := []struct {
		input    int
		expected bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d vs %v", tc.input, tc.expected), func(t *testing.T) {
			result := checkCriteria(tc.input)

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}
