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
			actual, _ := instructions(testCase.program)
			assertMemoryEquals(t, testCase.expected, actual)
		})
	}

	for _, testCase := range []struct {
		description string
		program     Memory
		expected    Memory
	}{
		{description: "1 + 1 = 2", program: Memory{1, 0, 0, 0, 99}, expected: Memory{2, 0, 0, 0, 99}},
		{description: "3 * 2 = 6", program: Memory{2, 3, 0, 3, 99}, expected: Memory{2, 3, 0, 6, 99}},
		{description: "99 * 99 = 9801", program: Memory{2, 4, 4, 5, 99, 0}, expected: Memory{2, 4, 4, 5, 99, 9801}},
		{description: "smoke test", program: Memory{1, 1, 1, 4, 99, 5, 6, 0, 99}, expected: Memory{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	} {
		t.Run(fmt.Sprintf("run %s", testCase.description), func(t *testing.T) {
			actual, _ := run(testCase.program)
			assertMemoryEquals(t, testCase.expected, actual)
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
			t.Errorf("Expected instructions %d to contain %d, but got %d", i, expected[i], actual[i])
		}
	}
}
