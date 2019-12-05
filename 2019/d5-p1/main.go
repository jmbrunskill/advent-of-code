package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Instruction struct {
	OpCode         int
	ParameterMode1 int
	ParameterMode2 int
	ParameterMode3 int
	ParameterMode4 int
	// InstPtrIncrement int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f))
}

func decodeInstruction(input int) Instruction {
	inst := Instruction{
		OpCode:         0,
		ParameterMode1: 0,
		ParameterMode2: 0,
		ParameterMode3: 0,
		ParameterMode4: 0,
		// InstPtrIncrement: 0,
	}

	inst.OpCode = input % 100
	input = input / 100

	inst.ParameterMode1 = input % 10
	input = input / 10

	inst.ParameterMode2 = input % 10
	input = input / 10

	inst.ParameterMode3 = input % 10
	input = input / 10

	inst.ParameterMode4 = input % 10
	input = input / 10

	return inst
}

func extractParam(input []int, location int, mode int) int {

	n := input[location]
	if mode == 0 {
		n = input[n]
	}

	return n
}

func runIntCode(input, code []int) ([]int, []int, error) {
	//Start at location 0 for the opcode
	instrucPtr := 0
	// fmt.Println("Processing input ", code, len(code))
	maxIterations := len(code) * 10 //Safeguard from infinite loop, I just assume 10 times the number of items in the input is a safe bet

	outputs := []int{}
	var inst Instruction
	var x int

	for i := 0; i < maxIterations; i++ {

		if len(code) <= instrucPtr {
			return nil, nil, fmt.Errorf("Error invalid instrucPtration %v", instrucPtr)
		}

		inst = decodeInstruction(code[instrucPtr])
		//First we need to check the op code
		switch inst.OpCode {
		case 1:
			n1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			n2 := extractParam(code, instrucPtr+2, inst.ParameterMode2)
			outLoc := code[instrucPtr+3]
			code[outLoc] = n1 + n2
			instrucPtr += 4 //move to next opcode
			// fmt.Printf("Encountered an addition operation at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 2:
			n1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			n2 := extractParam(code, instrucPtr+2, inst.ParameterMode2)
			outLoc := code[instrucPtr+3]
			code[outLoc] = n1 * n2
			instrucPtr += 4 //move to next opcode
			// fmt.Printf("Encountered an multiply operation at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 3:
			x, input = input[0], input[1:] //Pop off the input
			outLoc := code[instrucPtr+1]
			code[outLoc] = x
			instrucPtr += 2 //move to next opcode
			// fmt.Printf("Input at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 4:
			outLoc := code[instrucPtr+1]
			// fmt.Printf("OUTPUT %v\n", code[outLoc])
			outputs = append(outputs, code[outLoc])
			instrucPtr += 2 //move to next opcode
			// fmt.Printf("Output at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 99:
			//end of program
			// fmt.Println("Program Ended at Location:", instrucPtr)
			return code, outputs, nil
		default:
			return nil, nil, fmt.Errorf("Error invalid opcode %v at %d", code[instrucPtr], instrucPtr)
		}
		// fmt.Println(i, input)
	}

	return nil, nil, fmt.Errorf("Max Interations Reached %v", maxIterations)
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

	//Input a 1 as suggested for the diagnostics
	inputs := []int{1}

	//run intcode program to get output
	_, outputs, err := runIntCode(inputs, intProgramSlice)
	if err != nil {
		log.Fatalf("intcode program error %s", err)
	}

	return fmt.Sprintf("%v", outputs)
}
