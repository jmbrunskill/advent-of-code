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

	s := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`
	expected := "126"

	result := processInput(strings.NewReader(s))

	if result != expected {
		t.Fatalf("expected %v; got %v", expected, result)
	}
}
