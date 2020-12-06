package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidPassports(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{
			`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`, "2"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%v vs %v", tc.input, tc.expected), func(t *testing.T) {
			result := processInput(strings.NewReader(tc.input))

			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}
