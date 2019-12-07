package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLineSegments(t *testing.T) {
	tt := []struct {
		input    []string
		expected []lineSegment
	}{
		{[]string{"R8", "U5", "L5", "D3"}, []lineSegment{lineSegment{xy{0, 0}, xy{8, 0}}, lineSegment{xy{8, 0}, xy{8, 5}}, lineSegment{xy{8, 5}, xy{3, 5}}, lineSegment{xy{3, 5}, xy{3, 2}}}},
		{[]string{"U7", "R6", "D4", "L4"}, []lineSegment{lineSegment{xy{0, 0}, xy{0, 7}}, lineSegment{xy{0, 7}, xy{6, 7}}, lineSegment{xy{6, 7}, xy{6, 3}}, lineSegment{xy{6, 3}, xy{2, 3}}}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v -> %v", tc.input, tc.expected), func(t *testing.T) {
			result := lineSegments(tc.input)

			if !reflect.DeepEqual(tc.expected, result) {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tt := []struct {
		a        lineSegment
		b        lineSegment
		expected xy
	}{
		{lineSegment{xy{2, 3}, xy{6, 3}}, lineSegment{xy{3, 5}, xy{3, 2}}, xy{3, 3}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v intersect %v -> %v", tc.a, tc.b, tc.expected), func(t *testing.T) {
			result := lineIntersectionPoint(tc.a, tc.b)

			if !reflect.DeepEqual(tc.expected, *result) {
				t.Fatalf("expected %v; got %v", tc.expected, *result)
			}
		})
	}
}

func TestCrossDistance(t *testing.T) {
	tt := []struct {
		wire1    []string
		wire2    []string
		expected int
	}{

		{[]string{"R8", "U5", "L5", "D3"}, []string{"U7", "R6", "D4", "L4"}, 6},
		{[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}, []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}, 159},
		{[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}, []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}, 135},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v %v -> %v", tc.wire1, tc.wire2, tc.expected), func(t *testing.T) {
			result := calcCrossDistance(tc.wire1, tc.wire2)

			if tc.expected != result {
				t.Fatalf("expected %d; got %d", tc.expected, result)
			}
		})
	}
}
