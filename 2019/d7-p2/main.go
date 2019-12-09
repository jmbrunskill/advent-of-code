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

func extractParamLoc(code []int, location int, mode int, relativeBase int) ([]int, int) {
	// fmt.Println("extractParamLoc " /*code,*/, location, mode, relativeBase)

	l := location
	if mode == 0 {
		//Position mode
		code = memCheck(code, code[l])
		l = code[l]
	} else if mode == 1 {
		//Immediate Mode
		code = memCheck(code, l)
	} else if mode == 2 {
		// fmt.Println("extractParamLoc() processing ", location, mode, relativeBase, l)
		code = memCheck(code, relativeBase+l)
		l = relativeBase + code[l]
	} else {
		panic("Unimplemented parameter mode")
	}

	// fmt.Println("extractParamLoc returning", l)
	return code, l
}

func memCheck(mem []int, index int) []int {
	// fmt.Println("Memcheck(", index, ")", mem)
	if index < 0 {
		panic(fmt.Sprintf("Invalid mem index %v", index))
	}
	if index >= len(mem) {
		// fmt.Println("Memcheck(", index, ") increasing by", (index-len(mem))+1)
		//Increase slice to accept this index
		appendix := make([]int, (index-len(mem))+1)
		mem = append(mem, appendix...)
		// fmt.Println(len(mem))
	}
	return mem
}

func writeToMemory(mem []int, index, value int) []int {
	mem = memCheck(mem, index)
	mem[index] = value
	return mem
}

func runIntCode(inputs, code []int) ([]int, []int, error) {
	in := make(chan int, len(inputs)) //Buffer the inputs so we don't have to wait on them
	out := make(chan int)
	exitcode := make(chan int)
	errors := make(chan error, 0)

	go runIntCodeChans(code, in, out, exitcode, errors)

	//Send all the inputs to the channel
	for _, input := range inputs {
		fmt.Println("Input: ", input)
		in <- input
	}

	outputs := []int{}

	for {
		select {
		case v, ok := <-out:
			if !ok {
				// fmt.Println("ch", v, ok)
				return code, outputs, nil
			}
			outputs = append(outputs, v)
		case e := <-errors:
			// fmt.Println("Error", e)
			return code, outputs, e
		case v := <-exitcode:
			fmt.Println("Exited with ", v)
			return code, outputs, nil
		}

	}

	return code, outputs, nil
}

func runIntCodeChans(code []int, input, output, exitcode chan int, errors chan error) {
	//Start at location 0 for the opcode
	instrucPtr := 0

	//relative base
	relativeBase := 0

	// fmt.Printf("\n\n**\nProcessing code %v \n", code)
	maxIterations := len(code) * 10000 //Safeguard from infinite loop, I just assume 10 times the number of items in the input is a safe bet

	// outputs := []int{}
	var inst Instruction
	var x, p1_loc, p2_loc, p3_loc, lastOutput int

	for i := 0; i < maxIterations; i++ {

		if len(code) <= instrucPtr {
			errors <- fmt.Errorf("Error invalid instrucPtr %v", instrucPtr)
			return
		}

		inst = decodeInstruction(code[instrucPtr])
		// fmt.Printf("Instruction %v %v\n", code[instrucPtr], inst)
		//Extract Params
		if inst.OpCode != 99 {
			code, p1_loc = extractParamLoc(code, instrucPtr+1, inst.ParameterMode1, relativeBase)
		} else {
			p1_loc = -999
		}
		//OpCode Params for codes which require 2 params
		if inst.OpCode == 1 || inst.OpCode == 2 || inst.OpCode == 5 || inst.OpCode == 6 || inst.OpCode == 7 || inst.OpCode == 8 || inst.OpCode == 9 {
			code, p2_loc = extractParamLoc(code, instrucPtr+2, inst.ParameterMode2, relativeBase)
		} else {
			p2_loc = -999
		}
		//OpCode Params for codes which require 3 params
		if inst.OpCode == 1 || inst.OpCode == 2 || inst.OpCode == 7 || inst.OpCode == 8 {
			code, p3_loc = extractParamLoc(code, instrucPtr+3, inst.ParameterMode3, relativeBase)
		} else {
			p3_loc = -999
		}
		// fmt.Printf("Params at %v,%v,%v\n", p1_loc, p2_loc, p3_loc)

		//First we need to check the op code
		switch inst.OpCode {
		case 1:
			// fmt.Printf("Addition at %d %d * %d into %d\n", instrucPtr, code[p1_loc], code[p2_loc], p3_loc)
			//Makes sure we have the space for this parameter to be written too
			code = writeToMemory(code, p3_loc, code[p1_loc]+code[p2_loc])
			instrucPtr += 4 //move to next opcode
		case 2:
			// fmt.Printf("Multiply at %d %d * %d into %d = %d\n", instrucPtr, code[p1_loc], code[p2_loc], p3_loc, code[p1_loc]*code[p2_loc])
			code = writeToMemory(code, p3_loc, code[p1_loc]*code[p2_loc])
			instrucPtr += 4 //move to next opcode
		case 3:
			if inst.ParameterMode1 == 1 {
				errors <- fmt.Errorf("Error invalid Parameter Mode for instuction %v at %d", code[instrucPtr], instrucPtr)
				return
			}

			// fmt.Printf("Input at %d (%v)(%v)(%v) %d into %d\n", instrucPtr, code[instrucPtr], inst, code[instrucPtr+1], x, p1_loc)

			//Read input from the channel
			x = <-input

			code = writeToMemory(code, p1_loc, x)
			instrucPtr += 2 //move to next opcode
		case 4:
			// fmt.Printf("Output at %d - %v - %v\n", instrucPtr, code[instrucPtr], inst)
			// fmt.Printf("OUTPUT %v\n", code[p1_loc])
			// outputs = append(outputs, code[p1_loc])
			output <- code[p1_loc]
			lastOutput = code[p1_loc]
			instrucPtr += 2 //move to next opcode
		case 5:
			/*
				Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			*/
			// fmt.Printf("jump-if-true at %d - %d\n", instrucPtr, code[instrucPtr])
			// fmt.Printf("jump-if-true params %d, %d\n", code[p1_loc], code[p2_loc])
			if code[p1_loc] != 0 {
				instrucPtr = code[p2_loc]
			} else {
				instrucPtr += 3 //move to next opcode
			}
		case 6:
			/*
				Opcode 6 is jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			*/
			// fmt.Printf("jump-if-false at %d - %d\n", instrucPtr, code[instrucPtr])
			// fmt.Printf("jump-if-false params %d, %d\n", code[p1_loc], code[p2_loc])
			if code[p1_loc] == 0 {
				instrucPtr = code[p2_loc]
			} else {
				instrucPtr += 3 //move to next opcode
			}
		case 7:
			/*
				Opcode 7 is less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			*/
			// fmt.Printf("is-less-than at %d - %d\n", instrucPtr, code[instrucPtr])
			// fmt.Printf("is-equals params %d, %d, %d\n", code[p1_loc], code[p2_loc], p3_loc)
			if code[p1_loc] < code[p2_loc] {
				code = writeToMemory(code, p3_loc, 1)
			} else {
				code = writeToMemory(code, p3_loc, 0)
			}
			instrucPtr += 4 //move to next opcode
		case 8:
			/*
				Opcode 8 is equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			*/
			// fmt.Printf("is-equals at %d - %d\n", instrucPtr, code[instrucPtr])
			// fmt.Printf("is-equals params %d, %d, %d (%v)\n", code[p1_loc], code[p2_loc], p3_loc, inst)
			if code[p1_loc] == code[p2_loc] {
				code = writeToMemory(code, p3_loc, 1)
			} else {
				code = writeToMemory(code, p3_loc, 0)
			}
			instrucPtr += 4 //move to next opcode
		case 9:
			// fmt.Printf("Modify relative base at %d - %v - %v len(%v)\n", instrucPtr, code[instrucPtr], inst, len(code))
			relativeBase += code[p1_loc]
			// fmt.Printf("Relative base is now %d\n", relativeBase)
			instrucPtr += 2 //move to next opcode
		case 99:
			//end of program
			// fmt.Println("Program Ended at Location:", instrucPtr)
			exitcode <- lastOutput
			// close(input)
			close(exitcode)
			// return code, outputs, nil
			return
		default:
			errors <- fmt.Errorf("Error invalid opcode %v at %d", code[instrucPtr], instrucPtr)
			return
		}
		// fmt.Println(instrucPtr, code)
	}

	errors <- fmt.Errorf("Max Interations Reached %v", maxIterations)
	return
}

// func calcLoop(code []int, phase int, inputs, outputs chan) int {

// }

func calcSeq(code []int, seq []int) int {
	signal := 0

	fmt.Println("Calculating Seq", seq)

	chanA := make(chan int)
	chanB := make(chan int)
	chanC := make(chan int)
	chanD := make(chan int)
	chanE := make(chan int)
	exitA := make(chan int)
	exitB := make(chan int)
	exitC := make(chan int)
	exitD := make(chan int)
	exitE := make(chan int)
	errors := make(chan error, 0)

	//Copy code into new memory space
	memA := append([]int{}, code...)
	memB := append([]int{}, code...)
	memC := append([]int{}, code...)
	memD := append([]int{}, code...)
	memE := append([]int{}, code...)

	//Start all the AMPS running
	go runIntCodeChans(memA, chanE, chanA, exitA, errors)
	go runIntCodeChans(memB, chanA, chanB, exitB, errors)
	go runIntCodeChans(memC, chanB, chanC, exitC, errors)
	go runIntCodeChans(memD, chanC, chanD, exitD, errors)
	go runIntCodeChans(memE, chanD, chanE, exitE, errors)

	//Prime the engines

	//Channel seq Inputs for AMPs
	chanE <- seq[0]
	chanA <- seq[1]
	chanB <- seq[2]
	chanC <- seq[3]
	chanD <- seq[4]

	//Kick off the process by sending a 0 to AMP A (E's output)
	chanE <- 0

	//Wait for the programs to finish doing their thing

	select {
	case v, ok := <-exitA:
		if ok {
			fmt.Println("exitA Exited with ", v)
		}
	case v, ok := <-exitB:
		if ok {
			_ = v
			fmt.Println("B Exited with ", v)
		}
	case v, ok := <-exitC:
		if ok {
			fmt.Println("C Exited with ", v)
		}
	case v, ok := <-exitD:
		if ok {
			fmt.Println("D Exited with ", v)
			break //Cheating D should finish before E so we can safely break out of the loop here (ugly though)
		}
	case signal = <-exitE:
		fmt.Println("E Exited with the signal ", signal)
		break

	}

	//Get the final output from E
	signal = <-chanE
	fmt.Println("E ended with the signal ", signal)
	fmt.Printf("Returning Signal %v \n", signal)
	return signal
}

//Cheated a bit here https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
func generatePerms(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func findPhaseSeq(code []int) ([]int, int) {
	bestSeq := []int{}
	bestSignal := 0
	signal := 0

	combos := generatePerms([]int{9, 8, 7, 6, 5})

	for _, seq := range combos {
		signal = calcSeq(code, seq)

		if signal > bestSignal {
			bestSeq = seq
			bestSignal = signal
		}
	}

	return bestSeq, bestSignal
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

	//run intcode program to find phase output
	seq, signal := findPhaseSeq(intProgramSlice)

	return fmt.Sprintf("%v - %v", seq, signal)
}
