package main

import (
	"strings"
	"testing"
)

func TestOnesByTwos(t *testing.T) {

	s := `123456789012`
	result := processInput(strings.NewReader(s), 2, 3)
	expected := "1"
	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
