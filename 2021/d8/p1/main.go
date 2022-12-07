package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

func countUniqueOutputDigits(output []string) int {
	count := 0
	for _, v := range output {
		switch len(v) {
		case 2:
			//1 - 2 segments enabled
			count++
		case 4:
			//4 - 4 segments
			count++
		case 3:
			//7 - 3 segments
			count++
		case 7:
			//8 - 7 segments
			count++
		}
	}
	return count
}

func processInput(f io.Reader) string {
	startTime := time.Now()
	result := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		note := strings.Split(s.Text(), "|")
		// patterns := strings.Fields(note[0])
		output := strings.Fields(note[1])
		result += countUniqueOutputDigits(output)
		// fmt.Println(patterns, output)
	}

	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
