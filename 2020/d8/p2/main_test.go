package main

import (
	"strings"
	"testing"
)

func TestSimpleProgram(t *testing.T) {

	s := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	expected := "8"

	result := processInput(strings.NewReader(s))

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
