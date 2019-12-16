package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f, 100))
}

func getMultipliers(pos, needed int) []int {
	pattern := []int{0, 1, 0, -1}

	loc := 0
	repeats := 0

	var outputs = make([]int, 0, needed)
	for i := 0; i < needed; i++ {
		if repeats >= pos {
			repeats = 0
			loc++
			loc = loc % len(pattern)
		}
		outputs = append(outputs, pattern[loc])
		repeats++
	}
	return outputs
}

func processInput(f io.Reader, phases int) string {
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	inputs := []int{}
	for s.Scan() {
		// fmt.Println(s.Text())
		d, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Printf("%d\n", d)
		inputs = append(inputs, d)
	}

	outputs := []int{}
	for i := 0; i < phases; i++ {
		for j, _ := range inputs {
			mults := getMultipliers(j+1, len(inputs)+1)
			sum := 0
			for k := 0; k < len(inputs); k++ {
				sum += mults[k+1] * inputs[k]
			}
			outputs = append(outputs, abs(sum)%10)
		}

		// fmt.Println(outputs)
		inputs = outputs
		outputs = []int{}
	}
	str := ""
	for _, v := range inputs {
		str += strconv.Itoa(v)
	}

	return fmt.Sprintf("%v", str)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
