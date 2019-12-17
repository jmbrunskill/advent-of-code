package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestProcessInput(t *testing.T) {

	tt := []struct {
		input    string
		expected string
	}{
		{"03036732577212944063491565474664", "84462026"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v->%v", tc.input, tc.expected), func(t *testing.T) {
			// phases := 100
			phases := 1
			fft := processInput(strings.NewReader(tc.input), phases)
			if fft != tc.expected {
				t.Fatalf("expected %v; got %v", tc.expected, fft)
			}
		})
	}

}
func TestApplyMultipliers(t *testing.T) {
	tt := []struct {
		pos      int
		needed   int
		expected int
	}{
		{1, 8, 4},
		{2, 8, 8},
		{3, 8, 2},
		{4, 8, 2},
		{5, 8, 6},
		{6, 8, 1},
		{7, 8, 5},
		{8, 8, 8},
	}

	var inputs [650 * 10000]int
	for i := 0; i < 8; i++ {
		inputs[i] = i + 1
	}
	printSome(&inputs, 0)

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d,%d", tc.pos, tc.needed), func(t *testing.T) {
			result := applyMultipliers(tc.pos, tc.needed, &inputs)

			if !reflect.DeepEqual(tc.expected, result) {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}

}
