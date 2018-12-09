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
	var nums []int

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not decode line %s: %v", s.Text(), err)
		}
		nums = append(nums, n)
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	sum := 0
	sums := map[int]bool{0: true}
	for {
		//Loop forever - we'll break when we find a duplicated sum
		for _, n := range nums {
			sum += n
			if sums[sum] {
				//Found the first sum that is duplicated
				return fmt.Sprintf("%d", sum)
			}
			sums[sum] = true
		}
	}

}
