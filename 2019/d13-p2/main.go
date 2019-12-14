package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
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

func runIntCodeChans(code []int, output chan int, input chan int, inputRequest chan bool, exitcode chan int, errors chan error) {
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

			//Request an input
			inputRequest <- true

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

	//Insert the Coins...
	intProgramSlice[0] = 2

	var g gameState
	g.screen = make(map[xy]int)

	in := make(chan int)
	inputRequest := make(chan bool)
	out := make(chan int)
	exitcode := make(chan int)
	errors := make(chan error)

	go runIntCodeChans(intProgramSlice, out, in, inputRequest, exitcode, errors)

	go func() {
		//Handle the inputs
		for {
			select {
			case b, ok := <-inputRequest:
				if !ok {
					fmt.Println("InputRequest", b, ok)
					return
				}
				if b {
					//WARNING hideous solution, wait 1 millisecond for the game state to be updated before calculating paddle direction
					//Can't think of the right answer to fix this right now :(
					//Would have expected the state to be updated before this can run due to mutexes?
					//but for some reason it doesn't seem to be, or something else I don't understand is going on.
					time.Sleep(1 * time.Millisecond)
					in <- g.suggestPaddleDirection()
				} else {
					fmt.Println("Invalid Input Request")
					return
				}

			}
		}
	}()

	//Output instructions are x, y, id
	x := 0
	y := 0
	id := 0
	instCounter := 0
	outputCount := 0
	ballX := 0
	paddleX := 0

	for {
		select {
		case inst, ok := <-out:
			if !ok {
				fmt.Println("ch", inst, ok)
				return "NOTOK"
			}

			if instCounter == 0 {
				x = inst
				instCounter++
			} else if instCounter == 1 {
				y = inst
				instCounter++
			} else if instCounter == 2 {
				id = inst
				g.updateState(x, y, id)
				instCounter = 0
				if outputCount > 879 {
					//Only output the state when we've got the initial
					if ballX != g.ballX {
						// g.Print()
						ballX = g.ballX
					}
					if paddleX != g.paddleX {
						paddleX = g.paddleX
					}
				}
				outputCount++
			}

		case e := <-errors:
			fmt.Println("Int Code Error", e)
			panic("error")
		case _ = <-exitcode:
			// fmt.Println("Exited with ")
			g.Print()
			return fmt.Sprintf("Score: %v", g.score())
		}

	}

	return fmt.Sprintf("%v", g.score())
}

type gameState struct {
	screen  map[xy]int
	ballX   int
	paddleX int
	mux     sync.Mutex
}

func (g *gameState) updateState(x, y, id int) {
	g.mux.Lock()
	g.screen[xy{x, y}] = id

	if id == 3 {
		g.paddleX = x
	} else if id == 4 {
		g.ballX = x
		// g.Print()
	}
	g.mux.Unlock()
}
func (g *gameState) score() int {
	g.mux.Lock()
	defer g.mux.Unlock()
	return g.screen[xy{-1, 0}]
}

func (g *gameState) suggestPaddleDirection() int {
	g.mux.Lock()
	defer g.mux.Unlock()
	if g.ballX < g.paddleX {
		return -1
	} else if g.ballX > g.paddleX {
		return 1
	}

	return 0
}

func (g *gameState) Print() {
	g.mux.Lock()
	paddleX := 0
	ballX := 0
	for y := 0; y < 20; y++ {
		for x := 0; x < 44; x++ {
			switch g.screen[xy{x, y}] {
			case 0:
				fmt.Printf(" ")
			case 1:
				fmt.Printf("|")
			case 2:
				fmt.Printf("+")
			case 3:
				fmt.Printf("-")
				paddleX = x
			case 4:
				fmt.Printf("o")
				ballX = x

			}

		}
		fmt.Println()
	}
	g.mux.Unlock()
	fmt.Printf("SCORE:%d (ballX: %v, PaddleX %v) Direction: %d\n", g.score(), ballX, paddleX, g.suggestPaddleDirection())
}

type xy struct {
	x int
	y int
}

func (p xy) String() string {
	return fmt.Sprintf("<x=%d, y=%d>", p.x, p.y)
}
