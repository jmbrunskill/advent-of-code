package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func sumIn(preambleBuffer [25]int, preambleLength, sum int) bool {
	for i := 0; i < preambleLength; i++ {
		for j := i + 1; j < preambleLength; j++ {
			if preambleBuffer[i]+preambleBuffer[j] == sum {
				return true
			}
		}
	}

	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f, 25))
}

func processInput(f io.Reader, preambleLength int) string {
	startTime := time.Now()
	result := 0

	var preambleBuffer [25]int //25 is max premable length

	var i int
	var idx int

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		_, err := fmt.Sscanf(s.Text(), "%d", &i)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		idx++

		if idx > preambleLength && !sumIn(preambleBuffer, preambleLength, i) {
			result = i
			break
		}
		preambleBuffer[idx%preambleLength] = i
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)

	}

	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
