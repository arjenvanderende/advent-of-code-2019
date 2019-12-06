package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		mass     int
		expected int
	}{
		{mass: 12, expected: 2},
		{mass: 14, expected: 2},
		{mass: 1969, expected: 654},
		{mass: 100756, expected: 33583},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("requiredFuel with mass %d", testCase.mass), func(t *testing.T) {
			fuel := requiredFuel(testCase.mass)
			if fuel != testCase.expected {
				t.Error(fmt.Errorf("Expected %d fuel but got %d instead", testCase.expected, fuel))
			}
		})
	}
}
