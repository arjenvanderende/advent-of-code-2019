package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	for _, testCase := range []struct {
		program  string
		expected Memory
	}{
		{program: "1,0,0,0,99", expected: Memory{1, 0, 0, 0, 99}},
		{program: "2,3,0,3,99", expected: Memory{2, 3, 0, 3, 99}},
		{program: "2,4,4,5,99,0", expected: Memory{2, 4, 4, 5, 99, 0}},
	} {
		t.Run(fmt.Sprintf("instructions %s", testCase.program), func(t *testing.T) {
			actual, err := instructions(testCase.program)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			assertMemoryEquals(t, testCase.expected, actual)
		})
	}

	for _, testCase := range []struct {
		description string
		program     Memory
		input       int
		expected    Output
	}{
		// input
		{description: "return input (1)", program: Memory{3, 0, 4, 0, 99}, input: 1, expected: Output{1}},
		{description: "return input (50)", program: Memory{3, 0, 4, 0, 99}, input: 50, expected: Output{50}},

		// immediate / position mode
		{description: "multiply", program: Memory{1002, 4, 3, 4, 33}, input: 0, expected: Output{}},
		{description: "add with negative numbers", program: Memory{1101, 100, -1, 4, 0}, input: 0, expected: Output{}},

		// equals
		{description: "equals with position mode", program: Memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 7, expected: Output{0}},
		{description: "equals with position mode", program: Memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 8, expected: Output{1}},
		{description: "equals with position mode", program: Memory{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 9, expected: Output{0}},
		{description: "equals with immediate mode", program: Memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 7, expected: Output{0}},
		{description: "equals with immediate mode", program: Memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 8, expected: Output{1}},
		{description: "equals with immediate mode", program: Memory{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 9, expected: Output{0}},

		// less than
		{description: "less than with position mode", program: Memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: 7, expected: Output{1}},
		{description: "less than with position mode", program: Memory{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: 8, expected: Output{0}},
		{description: "less than with immediate mode", program: Memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: 7, expected: Output{1}},
		{description: "less than with immediate mode", program: Memory{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: 8, expected: Output{0}},

		// jump
		{description: "jump with position mode", program: Memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: 0, expected: Output{0}},
		{description: "jump with position mode", program: Memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: 1, expected: Output{1}},
		{description: "jump with position mode", program: Memory{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: 5, expected: Output{1}},
		{description: "jump with immediate mode", program: Memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 0, expected: Output{0}},
		{description: "jump with immediate mode", program: Memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 1, expected: Output{1}},
		{description: "jump with immediate mode", program: Memory{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 5, expected: Output{1}},

		// smoke test
		{description: "larger example", input: 7, expected: Output{999}, program: Memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
		{description: "larger example", input: 8, expected: Output{1000}, program: Memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
		{description: "larger example", input: 9, expected: Output{1001}, program: Memory{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
	} {
		t.Run(fmt.Sprintf("run %s", testCase.description), func(t *testing.T) {
			_, actual, err := run(testCase.program, testCase.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			assertOutputEquals(t, testCase.expected, actual)
		})
	}
}

func assertMemoryEquals(t *testing.T, expected Memory, actual Memory) {
	if len(actual) != len(expected) {
		t.Errorf("Expected %d instructions, but got %d", len(expected), len(actual))
		return
	}

	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Expected instruction %d to contain %d, but got %d", i, expected[i], actual[i])
		}
	}
}

func assertOutputEquals(t *testing.T, expected Output, actual Output) {
	if len(actual) != len(expected) {
		t.Errorf("Expected %d outputs, but got %d", len(expected), len(actual))
		return
	}

	for i := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Expected output %d to contain %d, but got %d", i, expected[i], actual[i])
		}
	}
}
