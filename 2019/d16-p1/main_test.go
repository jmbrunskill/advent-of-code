package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGetMultipliers(t *testing.T) {

	tt := []struct {
		pos      int
		needed   int
		expected []int
	}{
		{3, 12, []int{0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1}},
		{2, 12, []int{0, 0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d,%d", tc.pos, tc.needed), func(t *testing.T) {
			result := getMultipliers(tc.pos, tc.needed)

			if !reflect.DeepEqual(tc.expected, result) {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}

}

func TestSimpleCase(t *testing.T) {

	input := "12345678"
	phases := 4
	expected := "01029498"

	fft := processInput(strings.NewReader(input), phases)

	if fft != expected {
		t.Fatalf("expected %v; got %v", expected, fft)
	}
}
