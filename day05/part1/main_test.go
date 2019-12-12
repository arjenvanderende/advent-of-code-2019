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
		{description: "return input (1)", program: Memory{3, 0, 4, 0, 99}, input: 1, expected: Output{1}},
		{description: "return input (50)", program: Memory{3, 0, 4, 0, 99}, input: 50, expected: Output{50}},
		{description: "multiply mixed mode", program: Memory{1002, 4, 3, 4, 33}, input: 0, expected: Output{}},
		{description: "add mixed mode with negative", program: Memory{1101, 100, -1, 4, 0}, input: 0, expected: Output{}},
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
