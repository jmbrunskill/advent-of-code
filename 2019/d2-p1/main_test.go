package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntcode(t *testing.T) {
	tt := []struct {
		input    []int
		expected []int
		err      bool
	}{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, false},
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, false},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, false},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, false},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, false},
		{[]int{1, 0, 0, 0}, []int{}, true},     //error - No end instruction
		{[]int{7, 0, 0, 0, 99}, []int{}, true}, //error - invalid opcode
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.input, tc.expected), func(t *testing.T) {
			result, err := runIntCode(tc.input)

			// if err != nil {
			// 	fmt.Printf("%v -> %v (%v)\n", tc.input, tc.expected, err)
			// }

			if tc.err && err == nil {
				t.Fatalf("expected an error but got none %v", err)
			} else if !tc.err {
				if err != nil {
					t.Fatalf("expected no error but got %v", err)
				}
				if !reflect.DeepEqual(tc.expected, result) {
					t.Fatalf("expected %v; got %v", tc.expected, result)
				}
			}

		})
	}
}
