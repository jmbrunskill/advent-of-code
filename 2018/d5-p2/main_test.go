package main

import (
	"fmt"
	"testing"
)

func TestRection(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
		{"aA", ""},
		{"abBA", ""},
		{"aabAAB", "aabAAB"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf(tc.input, tc.input, tc.expected), func(t *testing.T) {

			str := react(tc.input)

			if tc.expected != str {
				t.Fatalf("expected %s; got %s", tc.expected, str)
			}
		})
	}
}

func TestRemoveUnits(t *testing.T) {
	tt := []struct {
		polymer  string
		unit     byte
		expected string
	}{
		{"aA", 'a', ""},
		{"abBA", 'a', "bB"},
		{"aabAAB", 'b', "aaAA"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf(tc.polymer, tc.polymer, tc.expected), func(t *testing.T) {

			str := removeUnits(tc.unit, tc.polymer)

			if tc.expected != str {
				t.Fatalf("expected %s; got %s", tc.expected, str)
			}
		})
	}
}
