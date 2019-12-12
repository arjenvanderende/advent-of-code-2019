package main

import (
	"fmt"
	"testing"
)

const orderedInput = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

const unorderedInput = `B)C
E)J
D)E
J)K
COM)B
D)I
E)F
B)G
C)D
G)H
K)L`

func TestMain(t *testing.T) {
	for _, testCase := range []struct {
		description string
		input       string
		expected    int
	}{
		{description: "ordered", input: orderedInput, expected: 42},
		{description: "unordered", input: unorderedInput, expected: 42},
	} {
		t.Run(fmt.Sprintf("orbits %s", testCase.description), func(t *testing.T) {
			actual, err := orbits(testCase.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if testCase.expected != actual {
				t.Errorf("expected %d got %d", testCase.expected, actual)
			}
		})
	}
}
