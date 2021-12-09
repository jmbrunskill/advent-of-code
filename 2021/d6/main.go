package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func simFish(initialFish []int, days int) int {

	//Inspired by Radix sort, count the number of fish at each generation and adjust accordingly each day
	var afishRadix [9]int //Thanks to go being awesome, this gets initialised with 0's
	var bfishRadix [9]int

	//Get the initial counts for the fish
	for i := 0; i < len(initialFish); i++ {
		afishRadix[initialFish[i]]++
	}

	/*
		Each day, a 0 becomes a 6 and adds a new 8 to the end of the list, while each other number decreases by 1 if it was present at the start of the day.
	*/
	currentFish := &afishRadix
	nextFish := &bfishRadix
	for d := 0; d < days; d++ {
		// fmt.Printf("Day %d : ", d)
		for i := 0; i < len(currentFish); i++ {
			// fmt.Printf("%d,", currentFish[i])
			switch i {
			case 0:
				nextFish[8] = currentFish[0]
			default:
				nextFish[i-1] = currentFish[i]
			}
		}
		nextFish[6] += currentFish[0]

		// fmt.Println()

		//Switch the arrays for calculation
		tmp := currentFish
		currentFish = nextFish
		nextFish = tmp

	}

	totalFish := 0
	//Sum up the fish
	for i := 0; i < len(currentFish); i++ {
		totalFish += currentFish[i]
	}

	return totalFish
}

func processInput(f io.Reader) string {
	startTime := time.Now()
	//Part1 numDays := 80 //Number of days to simulate fish growth for
	numDays := 256 //Number of days to simulate fish growth for

	initialFish := make([]int, 0)

	s := bufio.NewScanner(f)

	// Define a split function that separates on commas. (copied from https://golang.org/pkg/bufio/)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	s.Split(onComma)

	for s.Scan() {
		// fmt.Println(s.Text())
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		initialFish = append(initialFish, n)

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	result := simFish(initialFish, numDays)

	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))
	return fmt.Sprintf("%d", result)
}
