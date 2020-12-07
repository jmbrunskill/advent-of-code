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

	groupAnswers := make(map[rune]bool)

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		if s.Text() == "" {
			result += len(groupAnswers)
			//Reset the map so we can count the next group
			groupAnswers = make(map[rune]bool)
		} else {
			for _, c := range s.Text() {
				groupAnswers[c] = true
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}
	//Special case for the end of file
	result += len(groupAnswers)

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result %v in %d seconds\n", result, endTime-startTime)

	return fmt.Sprintf("%d", result)
}
