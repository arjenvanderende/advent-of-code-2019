package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	for _, testCase := range []struct {
		input    int
		expected bool
	}{
		// not 6 digits
		{input: 0, expected: false},
		{input: 99999, expected: false},
		{input: 1000000, expected: false},
		// decreasing pair of digits
		{input: 223450, expected: false},
		// no double
		{input: 123789, expected: false},
		// not all repeated digits are exactly two digits long
		{input: 111111, expected: false},
		{input: 123444, expected: false},
		// valid
		{input: 112233, expected: true},
		{input: 111122, expected: true},
	} {
		t.Run(fmt.Sprintf("valid %d", testCase.input), func(t *testing.T) {
			actual := valid(testCase.input)
			if testCase.expected != actual {
				t.Errorf("expected %v but got %v", testCase.expected, actual)
			}
		})
	}
}
