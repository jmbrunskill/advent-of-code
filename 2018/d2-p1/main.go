package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	numTwos := 0
	numThrees := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		letterCounts := map[rune]int{}
		foundTwos := 0
		foundThrees := 0
		for _, r := range s.Text() {
			letterCounts[r]++
			if letterCounts[r] == 2 {
				foundTwos++
			} else if letterCounts[r] == 3 {
				foundTwos--
				foundThrees++
			} else if letterCounts[r] > 3 {
				foundThrees--
			}
		}
		if foundTwos > 0 {
			numTwos++
		}
		if foundThrees > 0 {
			numThrees++
		}

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	return fmt.Sprintf("%d", numTwos*numThrees)

}
