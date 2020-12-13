package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f))
}

func processInput(f io.Reader) string {
	startTime := time.Now()
	result := 0

	adapters := make([]int, 0, 0)

	var joltage int

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		_, err := fmt.Sscanf(s.Text(), "%d", &joltage)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		adapters = append(adapters, joltage)
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)

	}

	sort.Ints(adapters)

	lastJoltage := 0

	diff1Count := 0
	diff2Count := 0
	diff3Count := 0

	for _, jolts := range adapters {
		switch jolts - lastJoltage {
		case 1:
			diff1Count++
		case 2:
			diff2Count++
		case 3:
			diff3Count++
		default:
			fmt.Printf("Error invalid jump from %v to %v\n", lastJoltage, jolts)
		}
		lastJoltage = jolts
	}

	fmt.Printf("1: %d, 2: %d, 3: %d\n", diff1Count, diff2Count, diff3Count)
	result = diff1Count * (diff3Count + 1) // Plus one for the 3 joltage, because the last adaptor is always a plus 3

	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
