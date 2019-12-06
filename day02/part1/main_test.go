package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	for _, testCase := range []struct {
		program  string
		expected []int
	}{
		{program: "1,0,0,0,99", expected: []int{1, 0, 0, 0, 99}},
		{program: "2,3,0,3,99", expected: []int{2, 3, 0, 3, 99}},
		{program: "2,4,4,5,99,0", expected: []int{2, 4, 4, 5, 99, 0}},
	} {
		t.Run(fmt.Sprintf("instructions %s", testCase.program), func(t *testing.T) {
			actual, _ := instructions(testCase.program)
			assertInstructionsEquals(t, testCase.expected, actual)
		})
	}

	for _, testCase := range []struct {
		description string
		program     []int
		expected    []int
	}{
		{description: "1 + 1 = 2", program: []int{1, 0, 0, 0, 99}, expected: []int{2, 0, 0, 0, 99}},
		{description: "3 * 2 = 6", program: []int{2, 3, 0, 3, 99}, expected: []int{2, 3, 0, 6, 99}},
		{description: "99 * 99 = 9801", program: []int{2, 4, 4, 5, 99, 0}, expected: []int{2, 4, 4, 5, 99, 9801}},
		{description: "smoke test", program: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	} {
		t.Run(fmt.Sprintf("run %s", testCase.description), func(t *testing.T) {
			actual := run(testCase.program)
			assertInstructionsEquals(t, testCase.expected, actual)
		})
	}
}

func assertInstructionsEquals(t *testing.T, expected []int, actual []int) {
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
