package main

import (
	"strings"
	"testing"
)

func TestSimpleInput(t *testing.T) {

	s := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`
	expected := "62"

	result := processInput(strings.NewReader(s), 5)

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
