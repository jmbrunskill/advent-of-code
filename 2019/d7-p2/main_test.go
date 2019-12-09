package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeOp(t *testing.T) {
	tt := []struct {
		input    int
		expected Instruction
	}{
		{2, Instruction{
			OpCode:         2,
			ParameterMode1: 0,
			ParameterMode2: 0,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
		{1, Instruction{
			OpCode:         1,
			ParameterMode1: 0,
			ParameterMode2: 0,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
		{1002, Instruction{
			OpCode:         2,
			ParameterMode1: 0,
			ParameterMode2: 1,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
		{111102, Instruction{
			OpCode:         2,
			ParameterMode1: 1,
			ParameterMode2: 1,
			ParameterMode3: 1,
			ParameterMode4: 1,
		}},
		{204, Instruction{
			OpCode:         4,
			ParameterMode1: 2,
			ParameterMode2: 0,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
		{2004, Instruction{
			OpCode:         4,
			ParameterMode1: 0,
			ParameterMode2: 2,
			ParameterMode3: 0,
			ParameterMode4: 0,
		}},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Decoding %v into %v", tc.input, tc.expected), func(t *testing.T) {
			result := decodeInstruction(tc.input)
			if tc.expected != result {
				t.Fatalf("expected %v; got %v", tc.expected, result)
			}
		})
	}
}

func TestIntcode(t *testing.T) {
	tt := []struct {
		inputs   []int
		outputs  []int
		code     []int
		expected []int
		err      bool
	}{
		{[]int{}, []int{}, []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, false},
		{[]int{}, []int{}, []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, false},
		{[]int{}, []int{}, []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, false},
		{[]int{}, []int{}, []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, false},
		{[]int{}, []int{}, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, false},
		{[]int{}, []int{}, []int{1002, 4, 3, 4, 33}, []int{1002, 4, 3, 4, 99}, false},      //Immediate mode
		{[]int{}, []int{}, []int{1101, 100, -1, 4, 0}, []int{1101, 100, -1, 4, 99}, false}, //Immediate mode
		{[]int{99}, []int{}, []int{3, 2, 0}, []int{3, 2, 99}, false},                       //Input mode
		{[]int{99}, []int{99}, []int{3, 4, 4, 4, 0}, []int{3, 4, 4, 4, 99}, false},         //Input & Output mode
		{[]int{}, []int{}, []int{1, 0, 0, 0}, []int{}, true},                               //error - No end instruction
		{[]int{}, []int{}, []int{98, 0, 0, 0, 99}, []int{}, true},                          //error - invalid opcode

		{[]int{8}, []int{1}, []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8}, false}, //consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
		{[]int{1}, []int{0}, []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8}, false}, //consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).

		{[]int{8}, []int{0}, []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8}, false}, //consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
		{[]int{1}, []int{1}, []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8}, false}, //consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).

		{[]int{0}, []int{0}, []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9}, false}, //Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero
		{[]int{1}, []int{1}, []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 1, 1, 1, 9}, false}, //Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero

		{[]int{0}, []int{0}, []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0}, false}, //Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero
		{[]int{1}, []int{1}, []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, []int{3, 3, 1105, 1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, false}, //Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero

		// //uses an input instruction to ask for a single number. The program will then output 999 if the input value is below 8, output 1000 if the input value is equal to 8, or output 1001 if the input value is greater than 8.
		{[]int{0}, []int{999},
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
				1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
				999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, false}, //Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero
		{[]int{8}, []int{1000}, []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 1000, 8, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, false}, //Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero
		{[]int{9}, []int{1001}, []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, []int{3, 21, 1008, 21,
			8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 1001, 9, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, false}, //Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero

	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("In:%v Expected:%v Outputs: %v", tc.code, tc.expected, tc.outputs), func(t *testing.T) {
			result, outputs, err := runIntCode(tc.inputs, tc.code)
			if tc.err && err == nil {
				fmt.Println(tc.code, tc.outputs)
				t.Fatalf("expected an error but got none %v", err)

			} else if !tc.err {
				if err != nil {
					t.Fatalf("expected no error but got %v", err)
				}
				if !reflect.DeepEqual(tc.expected, result) {
					t.Fatalf("expected %v; got %v (%v)", tc.expected, result, outputs)
				}
				if !reflect.DeepEqual(tc.outputs, outputs) {
					t.Fatalf("expected output %v; got %v", tc.outputs, outputs)
				}
			}

		})
	}
}

func TestIntcodeRelative(t *testing.T) {
	//Day 9 Tests
	tt := []struct {
		inputs   []int
		outputs  []int
		code     []int
		expected []int
		err      bool
	}{
		//Slighly dodgy editing this test case, because the slice we are using for the code is not the full slice after adding the channels
		// {[]int{}, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		// 	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
		{[]int{}, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, false},
		{[]int{}, []int{1219070632396864}, []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, []int{1102, 34915192, 34915192, 7, 4, 7, 99, 1219070632396864}, false},
		{[]int{}, []int{1125899906842624}, []int{104, 1125899906842624, 99}, []int{104, 1125899906842624, 99}, false},
		{[]int{99}, []int{}, []int{203, 2, 0}, []int{203, 2, 99}, false}, //Relative Input mode
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("In:%v Expected:%v Outputs: %v", tc.code, tc.expected, tc.outputs), func(t *testing.T) {
			result, outputs, err := runIntCode(tc.inputs, tc.code)

			// if err != nil {
			// 	fmt.Printf("%v -> %v (%v)\n", tc.code, tc.expected, err)
			// }

			if tc.err && err == nil {
				fmt.Println(tc.code, tc.outputs)
				t.Fatalf("expected an error but got none %v", err)

			} else if !tc.err {
				if err != nil {
					t.Fatalf("expected no error but got %v", err)
				}
				if !reflect.DeepEqual(tc.expected, result) {
					t.Fatalf("expected %v; got %v (%v)", tc.expected, result, outputs)
				}
				if !reflect.DeepEqual(tc.outputs, outputs) {
					t.Fatalf("expected output %v; got %v", tc.outputs, outputs)
				}
			}

		})
	}
}

func TestFeedbackLoop(t *testing.T) {
	//Day 7 Tests
	tt := []struct {
		code   []int
		seq    []int
		signal int
	}{
		{[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
			27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}, []int{9, 8, 7, 6, 5}, 139629729},
		{[]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
			-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
			53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}, []int{9, 7, 8, 5, 6}, 18216},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("In:%v Expected:%v : %v", tc.code, tc.seq, tc.signal), func(t *testing.T) {
			seq, signal := findPhaseSeq(tc.code)

			if !reflect.DeepEqual(tc.seq, seq) {
				t.Fatalf("expected %v; got %v (%v)", tc.seq, seq, signal)
			}
			if tc.signal != signal {
				t.Fatalf("expected output %v; got %v", tc.signal, signal)
			}
		})
	}
}
