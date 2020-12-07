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

func processInput(f io.Reader) string {
	startTime := time.Now().Unix()
	result := 0

	groupAnswers := make(map[rune]int)
	groupSize := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		if s.Text() == "" {
			numYes := 0
			for _, v := range groupAnswers {
				if v == groupSize {
					numYes++
				}
			}
			result += numYes
			//Reset the map so we can count the next group
			groupSize = 0
			groupAnswers = make(map[rune]int)
		} else {
			groupSize++
			for _, c := range s.Text() {
				groupAnswers[c]++
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}
	//Special case for the end of file

	numYes := 0
	for _, v := range groupAnswers {
		if v == groupSize {
			numYes++
		}
	}
	result += numYes

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result %v in %d seconds\n", result, endTime-startTime)

	return fmt.Sprintf("%d", result)
}
