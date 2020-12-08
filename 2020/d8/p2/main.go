package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Instruction struct {
	Name  string
	Param int
}

func runCode(code []Instruction) (int, error) {

	InsPtr := 0
	Acc := 0

	maxIterations := len(code) * 10000 //Safeguard from infinite loop

	seenInstructions := make(map[int]bool)

	var inst Instruction
	for i := 0; i < maxIterations && InsPtr < len(code); i++ {
		inst = code[InsPtr]
		// fmt.Println(InsPtr, inst)

		if seenInstructions[InsPtr] {
			return Acc, fmt.Errorf("Infinite Loop Detected %v at %v", inst, InsPtr)
		}
		seenInstructions[InsPtr] = true

		switch inst.Name {
		case "acc":
			// acc increases or decreases a single global value called the accumulator by the value given in the argument. For example, acc +7 would increase the accumulator by 7. The accumulator starts at 0. After an acc instruction, the instruction immediately below it is executed next.
			// fmt.Println("acc")
			Acc += inst.Param
			InsPtr++
		case "jmp":
			// jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from the jmp instruction; for example, jmp +2 would skip the next instruction, jmp +1 would continue to the instruction immediately below it, and jmp -20 would cause the instruction 20 lines above to be executed next.
			// fmt.Println("jmp")
			InsPtr += inst.Param //Jump to next instruction
		case "nop":
			// nop stands for No OPeration - it does nothing. The instruction immediately below it is executed next.
			// fmt.Println("nop")
			InsPtr++
		default:
			return Acc, fmt.Errorf("Error invalid istruction %v at %d", inst, InsPtr)
		}

	}

	if InsPtr == len(code) {
		return Acc, nil
	}

	return Acc, fmt.Errorf("Out of bounds %d", InsPtr)
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

func processInput(f io.Reader) string {
	startTime := time.Now().Unix()
	result := 0

	programSlice := make([]Instruction, 0, 5)

	var inst string
	var param int

	s := bufio.NewScanner(f)
	for s.Scan() {
		// fmt.Println(s.Text())

		_, err := fmt.Sscanf(s.Text(), "%s %d", &inst, &param)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		programSlice = append(programSlice, Instruction{inst, param})
	}
	if err := s.Err(); err != nil {
		log.Fatal("Scan() - ", err)

	}

	modified := false
	for i, inst := range programSlice {
		programSliceCopy := append([]Instruction{}, programSlice...)
		if inst.Name == "jmp" {
			programSliceCopy[i].Name = "nop"
			modified = true
		}
		if inst.Name == "nop" {
			programSliceCopy[i].Name = "jmp"
			modified = true
		}
		if modified {
			r, err := runCode(programSliceCopy)
			if err == nil {
				fmt.Println("Successfully Ran!")
				result = r
				break
			}
		}
		modified = false

	}

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result %v in %d seconds\n", result, endTime-startTime)

	return fmt.Sprintf("%d", result)
}
