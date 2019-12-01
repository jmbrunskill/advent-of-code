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
		workers       int
		expected      string
	}{
		{"example0.txt", 2, "258"},
		{"example1.txt", 2, "123"},
		{"example2.txt", 2, "124"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s vs %s", tc.inputFileName, tc.expected), func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", filepath.FromSlash(tc.inputFileName)))
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			str := processInput(f, tc.workers)

			if tc.expected != str {
				t.Fatalf("expected %s; got %s", tc.expected, str)
			}
		})
	}
}
