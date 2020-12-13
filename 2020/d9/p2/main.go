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

func sumContiguousSum(fullList []int, weakness int) int {
	incSum := 0
	sumMin := 0
	sumMax := 0
	for i := 0; i < len(fullList); i++ {
		sumMin = fullList[i]
		incSum = fullList[i]
		// fmt.Printf("Starting with sumMin %v - i:%d,incSum: %v, weakness %v\n", sumMin, i, incSum, weakness)
		for j := i + 1; j < len(fullList); j++ {
			if sumMin > fullList[j] {
				sumMin = fullList[j]
			} else if sumMax < fullList[j] {
				sumMax = fullList[j]
			}

			incSum += fullList[j]
			// fmt.Printf("i:%d,j:%d,sumMin %v, thisVal: %v, incSum: %v,weakness %v\n", i, j, sumMin, fullList[j], incSum, weakness)
			if incSum > weakness {
				//Contiguous sum is too big, start a new set to calculate for
				break
			} else if incSum == weakness {
				return (sumMin + sumMax)
			}
		}
		incSum = 0
	}

	return -1
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
	weakness := 0

	var preambleBuffer [25]int //25 is max premable length
	fullList := make([]int, 0)

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

		if weakness == 0 && idx > preambleLength && !sumIn(preambleBuffer, preambleLength, i) {
			weakness = i
		}
		preambleBuffer[idx%preambleLength] = i
		fullList = append(fullList, i)
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)

	}

	result := sumContiguousSum(fullList, weakness)

	fmt.Printf("Calculated result %v in %v\n", result, time.Since(startTime))

	return fmt.Sprintf("%d", result)
}
