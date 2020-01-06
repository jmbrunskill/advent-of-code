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
	fmt.Println(processInput(f, 100, 10000, false))
}

func applyMultipliers(pos, inputLength int, inputs *[651 * 10000]int, outputs *[651 * 10000]int, cusum *[651 * 10000]int) {
	// fmt.Printf("applyMultipliers(%d,%d,[])\n", pos, inputLength)
	sum := 0
	innerSum := 0

	onesStartLocation := pos - 1
	onesEndLocation := (2 * pos) - 2
	negative := false

	for onesEndLocation < inputLength {
		if onesStartLocation == 0 {
			innerSum = cusum[onesEndLocation]
		} else {
			innerSum = cusum[onesEndLocation] - cusum[onesStartLocation-1]
		}
		if negative {
			sum -= innerSum
		} else {
			sum += innerSum
		}
		onesStartLocation = onesEndLocation + pos + 1
		onesEndLocation = onesStartLocation + pos - 1
		negative = !negative
	}
	if onesStartLocation < inputLength && onesEndLocation >= inputLength {
		innerSum = cusum[inputLength-1] - cusum[onesStartLocation-1]
		if negative {
			sum -= innerSum
		} else {
			sum += innerSum
		}
	}

	// fmt.Println()
	outputs[pos-1] = abs(sum) % 10
	return
}

func processInput(f io.Reader, phases int, repeats int, offsetMode bool) string {
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

	outputOffset := 0
	var err error
	if offsetMode {
		offset := ""
		for i := 0; i < 7; i++ {
			offset += strconv.Itoa(inputs[i])
		}
		outputOffset, err = strconv.Atoi(offset)
		if err != nil {
			panic(offset)
		}
	}

	fmt.Println("Offset:", outputOffset)

	//always repeat the input 1000 times
	arrayLen := len(inputs) * repeats
	fmt.Println("Array Length:", arrayLen)

	var newInputs [651 * 10000]int
	var newOutputs [651 * 10000]int
	var cusum [651 * 10000]int
	for i := 0; i < arrayLen; i++ {
		newInputs[i] = inputs[i%len(inputs)]
	}
	// printSome(&newInputs, outputOffset)

	//SETUP CHANNELS
	calcRequests := make(chan calcRequest, 10)
	calcDone := make(chan int, 10)
	phaseDone := make(chan int, 0)
	//SETUP WORKERS
	//After adding the cusum array, the concurrency actually slows this down...
	for i := 0; i < 1; i++ {
		go asyncApplyMultipliers(calcRequests, calcDone)

	}

	pos := -1
	a := &newInputs
	b := &newOutputs
	tmp := &newInputs

	for i := 0; i < phases; i++ {
		startTime := time.Now()
		currentSum := 0
		for j := 0; j < arrayLen; j++ {
			//Create a cusum array
			currentSum += a[j]
			cusum[j] = currentSum
			// printSome(&cusum, 0)
		}

		go func() {
			for j := 0; j < arrayLen; j++ {
				//Wait for the calcs to be done (just count how many)
				pos = <-calcDone
				_ = pos
			}
			phaseDone <- 0
		}()
		for j := 0; j < arrayLen; j++ {
			// b[j] = applyMultipliers(j+1, arrayLen, a, b)
			calcRequests <- calcRequest{j + 1, arrayLen, a, b, &cusum}
		}
		pos = <-phaseDone
		_ = pos

		// printSome(b, outputOffset)
		// printLots(b, 0, arrayLen, len(inputs)*4*(i+1))

		//Swap the arrays over
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

type calcRequest struct {
	pos     int
	needed  int
	inputs  *[651 * 10000]int
	outputs *[651 * 10000]int
	cusum   *[651 * 10000]int
}

func asyncApplyMultipliers(calcs chan calcRequest, done chan int) {
	for {
		c, more := <-calcs
		if more {
			// fmt.Println("received job", c)
			applyMultipliers(c.pos, c.needed, c.inputs, c.outputs, c.cusum)
			done <- c.pos
		} else {
			// fmt.Println("Exiting")
			return
		}
	}
	// return
}

func printLots(input *[651 * 10000]int, offset, total, split int) {

	for i := offset; i < offset+total; i++ {
		if i%split == 0 {
			//Split the input on a period to help see patterns
			fmt.Printf("\ni:%03d:", i)
		}
		fmt.Printf("%d", input[i])
	}
	fmt.Println()
}

func printSome(input *[651 * 10000]int, offset int) {
	for i := 0 + offset; i < offset+8; i++ {
		fmt.Printf("%d,", input[i])
	}
	fmt.Println()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
