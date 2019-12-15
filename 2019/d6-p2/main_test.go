package main

import (
	"strings"
	"testing"
)

func TestOrbitToSanta(t *testing.T) {

	s := `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
	orbitTransfers := processInput(strings.NewReader(s))
	expected := "4"
	if orbitTransfers != expected {
		t.Fatalf("expected %v; got %v", expected, orbitTransfers)
	}
}
