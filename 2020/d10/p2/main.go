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

	waysToGetTo := make(map[int]int)

	//To start with we have 1 way to make a circuit
	waysToGetTo[0] = 1
	maxAdaptor := 0

	for _, jolts := range adapters {

		//The ways to join our adaptors is the sum of number of ways to join the adaptors to get the values up to 3 less than this
		waysToGetTo[jolts] = waysToGetTo[jolts-3] + waysToGetTo[jolts-2] + waysToGetTo[jolts-1]
		// fmt.Printf("Ways to get to %d is %d\n", jolts, waysToGetTo[jolts])
		maxAdaptor = jolts //Don't have to check if this is the max, as we have sorted values
	}

	result = waysToGetTo[maxAdaptor]

	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
