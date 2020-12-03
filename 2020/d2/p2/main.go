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

func parseLine(s string) (min, max int, c, pw string) {
	_, err := fmt.Sscanf(strings.Replace(strings.Replace(s, ":", " ", 1), "-", " ", 1), "%d %d %s %s", &min, &max, &c, &pw)
	if err != nil {
		log.Fatalf("could not read %s: %v", s, err)
	}
	// fmt.Printf("%d-%d %s %s\n", min, max, c, pw)

	return min, max, c, pw
}

func validatePassword(min, max int, c, pw string) bool {

	return (pw[min-1] == c[0]) != (pw[max-1] == c[0])
}

func processInput(f io.Reader) string {
	startTime := time.Now().Unix()

	result := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		valid := validatePassword(parseLine(s.Text()))
		if valid {
			result++
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result in %d seconds\n", endTime-startTime)

	return fmt.Sprintf("%d", result)
}
