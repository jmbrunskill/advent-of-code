package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseBag(t *testing.T) {
	tt := []struct {
		input    string
		expected Bag
	}{
		{"light red bags contain 1 bright white bag, 2 muted yellow bags.", Bag{
			Name:               "light red",
			ContainedBagNames:  []string{"bright white", "muted yellow"},
			ContainedBagCounts: []int{1, 2},
		}},
		{"bright white bags contain 1 shiny gold bag.", Bag{
			Name:               "bright white",
			ContainedBagNames:  []string{"shiny gold"},
			ContainedBagCounts: []int{1},
		}},
		{"faded blue bags contain no other bags", Bag{
			Name: "faded blue",
		}},
		{"dotted black bags contain no other bags.", Bag{
			Name: "dotted black",
		}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Bag:%v", tc.input), func(t *testing.T) {
			result := parseBag(tc.input)
			if !reflect.DeepEqual(&tc.expected, result) {
				t.Fatalf("expected %v; got %v", &tc.expected, result)
			}
		})
	}
}

func TestShinyGoldBags(t *testing.T) {

	s := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	expected := "4"

	result := processInput(strings.NewReader(s))

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
