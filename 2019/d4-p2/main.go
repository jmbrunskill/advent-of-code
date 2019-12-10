package main

import (
	"fmt"
	"log"
	"strconv"
)

func getDigits(i int) []int {
	//String Method :(

	digits := []int{}

	str := fmt.Sprintf("%d", i)
	for _, char := range str {
		i, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal("Invalid Digit")
		}
		digits = append(digits, i)
	}

	return digits
}

func checkCriteria(i int) bool {
	//Check min/max
	if i < 100000 {
		return false
	}
	if i > 685979 {
		return false
	}
	digits := getDigits(i)

	ajacent := false
	// ajacentDigit := 0
	ajacentCount := 0
	lastDigit := 0

	// fmt.Println(i, digits)
	for _, d := range digits {
		//Increasing Digits
		if d < lastDigit {
			//Decreasing Digit
			return false
		}
		//Check Ajacent Digits
		if d == lastDigit {
			// fmt.Printf("Same Digit Repeated\n")
			ajacentCount++
		} else if ajacentCount == 1 {
			// fmt.Printf("Ajacent Count == 1\n")
			ajacent = true
		} else {
			// fmt.Printf("Different Digit\n")
			ajacentCount = 0
		}
		// fmt.Printf(" %v, %v, %v, %v\n", ajacent, ajacentCount, lastDigit, d)

		lastDigit = d
	}
	if ajacentCount == 1 {
		return true
	}

	return ajacent
}

func main() {
	min := 136818
	max := 685979

	numPossible := 0

	for i := min; i < max; i++ {
		if checkCriteria(i) {
			numPossible++
		}
	}

	fmt.Println(numPossible)

}
