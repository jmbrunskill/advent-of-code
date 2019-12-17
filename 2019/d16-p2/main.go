package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
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

func applyMultipliers(pos, needed int, inputs *[650 * 10000]int) int {
	// fmt.Printf("applyMultipliers(%d,%d,[])\n", pos, needed)
	sum := 0
	loc := 0
	repeats := 1

	for i := 0; i < needed; i++ {
		// fmt.Printf("L i %d, repeats %d, pos %d, loc %d\n", i, repeats, pos, loc)
		if repeats >= pos {
			repeats = 0
			loc++
			loc = loc % 4 //len(pattern)
			// fmt.Printf("New Location: i %d, repeats %d, pos %d, loc %d\n", i, repeats, pos, loc)
		}

		switch loc {
		case 1:
			//* by 1
			// fmt.Printf("+ %d * %d (%v,%v)\n", inputs[i], pattern[loc], loc, i)
			sum += inputs[i]
		case 3:
			//* by -1
			// fmt.Printf("+ %d * %d (%v,%v)\n", inputs[i], pattern[loc], loc, i)
			sum -= inputs[i]
		default:
			//This is a zero, we can skip ahead as there's no point adding up zeros
			// skip := (pos - repeats - 1)
			i += (pos - repeats - 1)
			repeats = (pos - 1) //This is about to be incremented...
			// fmt.Printf("i:%d loc:%d Skipping Ahead %d pos %d repeats %d\n", i, loc, skip, pos, repeats)
		}
		repeats++
	}
	// fmt.Println()
	return abs(sum) % 10
}

func processInput(f io.Reader, phases int) string {
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	sT := time.Now()

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

	offset := ""
	for i := 0; i < 7; i++ {
		offset += strconv.Itoa(inputs[i])
	}

	outputOffset, err := strconv.Atoi(offset)
	if err != nil {
		panic(offset)
	}
	fmt.Println("Offset:", outputOffset)

	//always repeat the input 1000 times
	arrayLen := len(inputs) * 10000
	fmt.Println("Array Length:", arrayLen)

	var newInputs [650 * 10000]int
	var newOutputs [650 * 10000]int
	for i := 0; i < arrayLen; i++ {
		newInputs[i] = inputs[i%len(inputs)]
	}
	printSome(&newInputs, outputOffset)

	a := &newInputs
	b := &newOutputs
	tmp := &newInputs

	for i := 0; i < phases; i++ {
		startTime := time.Now()
		for j := 0; j < arrayLen; j++ {
			b[j] = applyMultipliers(j+1, arrayLen, a)
		}
		// fmt.Print(".")
		//Swap the arrays over
		printSome(b, outputOffset)
		tmp = a
		a = b
		b = tmp
		t := time.Now()
		fmt.Printf("Calculated phase %d in %v\n", i, t.Sub(startTime))
	}

	str := ""
	//Get the 8 digits
	for j := outputOffset; j < outputOffset+8; j++ {
		str += strconv.Itoa(a[j])
	}

	t := time.Now()
	fmt.Printf("Calculated result for %d in %v\n", phases, t.Sub(sT))
	return fmt.Sprintf("%v", str)
}

func printSome(input *[650 * 10000]int, offset int) {
	for i := 0 + offset; i < offset+10; i++ {
		fmt.Printf("%d", input[i])
	}
	fmt.Println()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
