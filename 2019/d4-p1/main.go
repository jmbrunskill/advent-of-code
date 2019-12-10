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
	if i < 136818 {
		return false
	}
	if i > 685979 {
		return false
	}
	digits := getDigits(i)

	ajacent := false
	lastDigit := 0

	for _, d := range digits {
		//Increasing Digits
		if d < lastDigit {
			//Decreasing Digit
			return false
		}
		//Check Ajacent Digits
		if d == lastDigit {
			ajacent = true
		}
		lastDigit = d
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
