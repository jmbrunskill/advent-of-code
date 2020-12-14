package main

import (
	"strings"
	"testing"
)

func TestSimpleExample(t *testing.T) {

	s := `939
7,13,x,x,59,x,31,19`
	expected := "295"

	result := processInput(strings.NewReader(s))

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
