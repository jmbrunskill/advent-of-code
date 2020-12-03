package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestTreeCount(t *testing.T) {
	tt := []struct {
		input    string
		down     int
		right    int
		expected string
	}{
		{
			`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`, 1, 3, "7"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.input, tc.expected), func(t *testing.T) {
			result := processInput(strings.NewReader(tc.input), tc.down, tc.right)

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}
