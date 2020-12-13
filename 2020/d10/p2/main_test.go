package main

import (
	"strings"
	"testing"
)

func TestSmallJolts(t *testing.T) {

	s := `16
10
15
5
1
11
7
19
6
12
4`
	expected := "8"

	result := processInput(strings.NewReader(s))

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}

func TestBiggerJolts(t *testing.T) {

	s := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`
	expected := "19208"

	result := processInput(strings.NewReader(s))

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
