package main

import (
	"strings"
	"testing"
)

func TestOrbitInfo(t *testing.T) {

	s := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	numOrbits := processInput(strings.NewReader(s))

	if numOrbits != "42" {
		t.Fatalf("expected 42; got %v", numOrbits)
	}
}

func TestOrbitInfoReorderd(t *testing.T) {

	s := `
B)C
COM)B
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
B)M`
	numOrbits := processInput(strings.NewReader(s))
	expected := "44"
	if numOrbits != expected {
		t.Fatalf("expected %v; got %v", expected, numOrbits)
	}
}
