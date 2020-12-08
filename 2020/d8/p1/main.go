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

type Console struct {
	Acc    int
	InsPtr int
	Code   []Instruction
}

func (c *Console) runTillRepeat() int {

	maxIterations := len(c.Code) * 10000 //Safeguard from infinite loop

	seenInstructions := make(map[int]bool)

	var inst Instruction
	for i := 0; i < maxIterations; i++ {
		inst = c.Code[c.InsPtr]
		fmt.Println(c.InsPtr, inst)

		if seenInstructions[c.InsPtr] {
			return c.Acc
		}
		seenInstructions[c.InsPtr] = true

		switch inst.Name {
		case "acc":
			// fmt.Println("acc")
			c.Acc += inst.Param
			c.InsPtr += 1
			// acc increases or decreases a single global value called the accumulator by the value given in the argument. For example, acc +7 would increase the accumulator by 7. The accumulator starts at 0. After an acc instruction, the instruction immediately below it is executed next.
		case "jmp":
			// fmt.Println("jmp")
			// jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from the jmp instruction; for example, jmp +2 would skip the next instruction, jmp +1 would continue to the instruction immediately below it, and jmp -20 would cause the instruction 20 lines above to be executed next.
			c.InsPtr += inst.Param //Jump to next instruction
		case "nop":
			// fmt.Println("nop")
			// nop stands for No OPeration - it does nothing. The instruction immediately below it is executed next.
			c.InsPtr += 1
		default:
			fmt.Printf("Error invalid istruction %v at %d", inst, c.InsPtr)
			return -1
		}

	}

	return -1
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

	console := Console{0, 0, programSlice}
	result = console.runTillRepeat()

	endTime := time.Now().Unix()
	fmt.Printf("Calculated result %v in %d seconds\n", result, endTime-startTime)

	return fmt.Sprintf("%d", result)
}
