package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestOldProcessInput(t *testing.T) {

	tt := []struct {
		input    string
		phases   int
		expected string
	}{
		{"12345678", 4, "01029498"},
		{"80871224585914546619083218645595", 100, "24176176"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v->%v", tc.input, tc.expected), func(t *testing.T) {
			fft := processInput(strings.NewReader(tc.input), tc.phases, 1, false)
			if fft != tc.expected {
				t.Fatalf("expected %v; got %v", tc.expected, fft)
			}
		})
	}

}
func TestProcessInput(t *testing.T) {

	tt := []struct {
		input    string
		phases   int
		repeats  int
		expected string
	}{
		// {"03036732577212944063491565474664", 1, 10000, "61706040"},
		{"03036732577212944063491565474664", 100, 10000, "84462026"},
		{"02935109699940807407585447034323", 100, 10000, "78725270"},
		{"03081770884921959731165446850517", 100, 10000, "53553731"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v->%v", tc.input, tc.expected), func(t *testing.T) {
			fft := processInput(strings.NewReader(tc.input), tc.phases, tc.repeats, true)
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

	arrayLen := 8
	var inputs [650 * 10000]int
	var outputs [650 * 10000]int
	for i := 0; i < arrayLen; i++ {
		inputs[i] = i + 1
		// fmt.Printf("%d,", inputs[i])
	}

	var cusum [650 * 10000]int
	currentSum := 0
	// fmt.Println("CUSUM:")
	for j := 0; j < arrayLen; j++ {
		//Create a cusum array
		currentSum += inputs[j]
		cusum[j] = currentSum
		// fmt.Printf("%d,", cusum[j])
	}
	// fmt.Println()

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d,%d", tc.pos, tc.needed), func(t *testing.T) {
			applyMultipliers(tc.pos, tc.needed, &inputs, &outputs, &cusum)

			if !reflect.DeepEqual(tc.expected, outputs[tc.pos-1]) {
				t.Fatalf("expected %v; got %v", tc.expected, outputs[tc.pos-1])
			}
		})
	}

}
