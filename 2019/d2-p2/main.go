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

func runIntCode(input []int) ([]int, error) {
	//Start at location 0 for the opcode
	instrucPtr := 0
	// fmt.Println("Processing input ", input, len(input))
	maxIterations := len(input) * 10 //Safeguard from infinite loop, I just assume 10 times the number of items in the input is a safe bet

	for i := 0; i < maxIterations; i++ {

		if len(input) <= instrucPtr {
			return nil, fmt.Errorf("Error invalid instrucPtr %v", instrucPtr)
		}

		//First we need to check the op code
		switch input[instrucPtr] {
		case 1:
			l1 := input[instrucPtr+1]
			l2 := input[instrucPtr+2]
			n1 := input[l1]
			n2 := input[l2]
			outLoc := input[instrucPtr+3]
			input[outLoc] = n1 + n2
			instrucPtr += 4 //move to next opcode
			// fmt.Printf("Encountered an addition operation at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 2:
			l1 := input[instrucPtr+1]
			l2 := input[instrucPtr+2]
			n1 := input[l1]
			n2 := input[l2]
			outLoc := input[instrucPtr+3]
			input[outLoc] = n1 * n2
			instrucPtr += 4 //move to next opcode
			// fmt.Printf("Encountered an multiply operation at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 99:
			//end of program
			// fmt.Println("Program Ended at Location:", instrucPtr)
			return input, nil
		default:
			return nil, fmt.Errorf("Error invalid opcode %v at %d", input[instrucPtr], instrucPtr)
		}
		// fmt.Println(i, input)

	}

	return nil, fmt.Errorf("Max Interations Reached %v", maxIterations)
}

func processInput(f io.Reader) string {
	s := bufio.NewScanner(f)

	intProgramSlice := make([]int, 0, 5)

	// Define a split function that separates on commas. (copied from https://golang.org/pkg/bufio/)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	s.Split(onComma)

	for s.Scan() {
		// fmt.Println(s.Text())
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		intProgramSlice = append(intProgramSlice, n)

	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)
	}

	//Try to find the inputs that create the output 19690720 (exhustive search)
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			//Copy the input into a fresh slice
			workingSlice := append([]int{}, intProgramSlice...)

			workingSlice[1] = i
			workingSlice[2] = j

			//run intcode program to get output
			result, err := runIntCode(workingSlice)
			if err != nil {
				log.Fatalf("intcode program error %s", err)
			}

			if result[0] == 19690720 {
				return fmt.Sprintf("%d", 100*i+j)
			}
			// fmt.Println(result)
		}
	}

	return fmt.Sprintf("NO RESULT")
}
