package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseReaction(t *testing.T) {
	tt := []struct {
		input    string
		expected reaction
	}{{"2 AB, 3 BC, 4 CA => 1 FUEL", reaction{
		name:        "FUEL",
		outputCount: 1,
		inputCounts: []int{2, 3, 4},
		inputNames:  []string{"AB", "BC", "CA"},
	}}}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Reaction:%v", tc.input), func(t *testing.T) {
			result := parseReaction(tc.input)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}

func TestSimpleReaction(t *testing.T) {

	s := `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`
	expected := "31"

	oreUsed := processInput(strings.NewReader(s))

	if oreUsed != expected {
		t.Fatalf("expected %v; got %v", expected, oreUsed)
	}
}

func TestComplexReaction(t *testing.T) {

	s := `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`
	expected := "2210736"

	oreUsed := processInput(strings.NewReader(s))

	if oreUsed != expected {
		t.Fatalf("expected %v; got %v", expected, oreUsed)
	}
}
