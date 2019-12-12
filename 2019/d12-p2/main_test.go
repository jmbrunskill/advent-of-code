package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculateGravity(t *testing.T) {
	tt := []struct {
		m1          moon
		m2          moon
		expected_m1 moon
		expected_m2 moon
	}{
		{moon{xyz{3, 0, 0}, xyz{0, 0, 0}}, moon{xyz{5, 0, 0}, xyz{0, 0, 0}}, moon{xyz{3, 0, 0}, xyz{1, 0, 0}}, moon{xyz{5, 0, 0}, xyz{-1, 0, 0}}},
		{moon{xyz{0, 3, 0}, xyz{0, 0, 0}}, moon{xyz{0, 5, 0}, xyz{0, 0, 0}}, moon{xyz{0, 3, 0}, xyz{0, 1, 0}}, moon{xyz{0, 5, 0}, xyz{0, -1, 0}}},
		{moon{xyz{0, 0, 3}, xyz{0, 0, 0}}, moon{xyz{0, 0, 5}, xyz{0, 0, 0}}, moon{xyz{0, 0, 3}, xyz{0, 0, 1}}, moon{xyz{0, 0, 5}, xyz{0, 0, -1}}},
		{moon{xyz{0, 0, 0}, xyz{0, 0, 0}}, moon{xyz{0, 0, 0}, xyz{0, 0, 0}}, moon{xyz{0, 0, 0}, xyz{0, 0, 0}}, moon{xyz{0, 0, 0}, xyz{0, 0, 0}}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v, %v -> %v,%v", tc.m1, tc.m2, tc.expected_m1, tc.expected_m2), func(t *testing.T) {
			m1, m2 := calculateGravity(tc.m1, tc.m2)

			if !reflect.DeepEqual(tc.expected_m1, m1) {
				t.Fatalf("expected %v; got %v", tc.expected_m1, m1)
			}
			if !reflect.DeepEqual(tc.expected_m2, m2) {
				t.Fatalf("expected %v; got %v", tc.expected_m2, m2)
			}
		})
	}
}

func TestCalculateEnergy(t *testing.T) {
	tt := []struct {
		m        moon
		expected int
	}{
		{moon{xyz{2, 1, 3}, xyz{3, 2, 1}}, 36},
		{moon{xyz{1, 8, 0}, xyz{-1, 1, 3}}, 45},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d -> %v", tc.m, tc.expected), func(t *testing.T) {
			result := calculateEnergy(tc.m)

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}

func TestRepeatingSteps(t *testing.T) {
	tt := []struct {
		moons    []xyz
		expected int
	}{
		{[]xyz{xyz{-1, 0, 2}, xyz{2, -10, -7}, xyz{4, -8, 8}, xyz{3, 5, -1}}, 2772},
		{[]xyz{xyz{-8, -10, 0}, xyz{5, 5, 10}, xyz{2, -7, 3}, xyz{9, -8, -3}}, 4686774924},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d -> %v", tc.moons, tc.expected), func(t *testing.T) {
			result := calculateRepeatSteps(tc.moons)

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}
