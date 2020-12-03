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

func mult2020_3(expenses []int) int {

	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return expenses[i] * expenses[j] * expenses[k]
				}
			}
		}
	}

	return 0
}

func processInput(f io.Reader) string {
	startTime := time.Now()
	expenses := make([]int, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		expenses = append(expenses, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	result := mult2020_3(expenses)

	fmt.Printf("Calculated result in %s \n", time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
