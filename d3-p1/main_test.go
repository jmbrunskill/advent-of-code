package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestProcessInput(t *testing.T) {
	tt := []struct {
		inputFileName string
		expected      string
	}{
		{"example0.txt", "0"},
		{"example1.txt", "4"},
		{"example2.txt", "6"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s,%s", tc.inputFileName, tc.expected), func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", filepath.FromSlash(tc.inputFileName)))
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			str := processInput(f)

			if tc.expected != str {
				t.Fatalf("expected %s; got %s", tc.expected, str)
			}
		})
	}
}
