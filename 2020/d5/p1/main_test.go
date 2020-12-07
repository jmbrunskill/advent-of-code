package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSeatId(t *testing.T) {
	tt := []struct {
		input    string
		expected int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.input, tc.expected), func(t *testing.T) {
			result := seatId(tc.input)

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}

func TestHighestSeat(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{
			`BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`, "820"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.input, tc.expected), func(t *testing.T) {
			result := processInput(strings.NewReader(tc.input))

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}
