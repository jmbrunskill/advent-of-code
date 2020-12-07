package main

import (
	"strings"
	"testing"
)

func TestYesQs(t *testing.T) {

	s := `abc

a
b
c

ab
ac

a
a
a
a

b`
	expected := "11"

	result := processInput(strings.NewReader(s))

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
