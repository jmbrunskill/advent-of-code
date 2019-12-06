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
	// fmt.Printf("Processing input %v %v\n", code, input)
	maxIterations := len(code) * 10 //Safeguard from infinite loop, I just assume 10 times the number of items in the input is a safe bet

	outputs := []int{}
	var inst Instruction
	var x int

	for i := 0; i < maxIterations; i++ {

		if len(code) <= instrucPtr {
			return nil, nil, fmt.Errorf("Error invalid instrucPtr %v", instrucPtr)
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
			// fmt.Printf("Addition at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 2:
			n1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			n2 := extractParam(code, instrucPtr+2, inst.ParameterMode2)
			outLoc := code[instrucPtr+3]
			code[outLoc] = n1 * n2
			instrucPtr += 4 //move to next opcode
			// fmt.Printf("Multiply at %d %d * %d into %d\n", instrucPtr, n1, n2, outLoc)
		case 3:
			x, input = input[0], input[1:] //Pop off the input
			outLoc := code[instrucPtr+1]
			code[outLoc] = x
			// fmt.Printf("Input at %d into %d\n", instrucPtr, outLoc)
			instrucPtr += 2 //move to next opcode
		case 4:
			// fmt.Printf("Output at %d - %v - %v len(%v)\n", instrucPtr, code[instrucPtr], inst, len(code))
			p1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			// fmt.Printf("OUTPUT %v\n", p1)
			outputs = append(outputs, p1)
			instrucPtr += 2 //move to next opcode
		case 5:
			/*
				Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			*/
			// fmt.Printf("jump-if-true at %d - %d\n", instrucPtr, code[instrucPtr])
			p1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			p2 := extractParam(code, instrucPtr+2, inst.ParameterMode2)
			// fmt.Printf("jump-if-true params %d, %d\n", p1, p2)
			if p1 != 0 {
				instrucPtr = p2
			} else {
				instrucPtr += 3 //move to next opcode
			}
		case 6:
			/*
				Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			*/
			// fmt.Printf("jump-if-false at %d - %d\n", instrucPtr, code[instrucPtr])
			p1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			p2 := extractParam(code, instrucPtr+2, inst.ParameterMode2)
			// fmt.Printf("jump-if-false params %d, %d\n", p1, p2)
			if p1 == 0 {
				instrucPtr = p2
			} else {
				instrucPtr += 3 //move to next opcode
			}
		case 7:
			/*
				Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			*/
			// fmt.Printf("is-less-than at %d - %d\n", instrucPtr, code[instrucPtr])
			p1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			p2 := extractParam(code, instrucPtr+2, inst.ParameterMode2)
			outLoc := code[instrucPtr+3]
			// fmt.Printf("is-equals params %d, %d, %d\n", p1, p2, outLoc)
			if p1 < p2 {
				code[outLoc] = 1
			} else {
				code[outLoc] = 0
			}
			instrucPtr += 4 //move to next opcode
		case 8:
			/*
				Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			*/
			// fmt.Printf("is-equals at %d - %d\n", instrucPtr, code[instrucPtr])
			p1 := extractParam(code, instrucPtr+1, inst.ParameterMode1)
			p2 := extractParam(code, instrucPtr+2, inst.ParameterMode2)
			outLoc := code[instrucPtr+3]
			// fmt.Printf("is-equals params %d, %d, %d (%v)\n", p1, p2, outLoc, inst)
			if p1 == p2 {
				code[outLoc] = 1
			} else {
				code[outLoc] = 0
			}
			instrucPtr += 4 //move to next opcode
		case 99:
			//end of program
			// fmt.Println("Program Ended at Location:", instrucPtr)
			return code, outputs, nil
		default:
			return nil, nil, fmt.Errorf("Error invalid opcode %v at %d", code[instrucPtr], instrucPtr)
		}
		// fmt.Println(instrucPtr, code)
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

	//Input 5 for the Radiator Controller
	inputs := []int{5}

	//run intcode program to get output
	_, outputs, err := runIntCode(inputs, intProgramSlice)
	if err != nil {
		log.Fatalf("intcode program error %s", err)
	}

	return fmt.Sprintf("%v", outputs)
}
