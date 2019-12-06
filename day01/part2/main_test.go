package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	for _, testCase := range []struct {
		mass     int
		expected int
	}{
		{mass: 12, expected: 2},
		{mass: 14, expected: 2},
		{mass: 1969, expected: 654},
		{mass: 100756, expected: 33583},
	} {
		t.Run(fmt.Sprintf("requiredFuel with mass %d", testCase.mass), func(t *testing.T) {
			fuel := requiredFuel(testCase.mass)
			if fuel != testCase.expected {
				t.Error(fmt.Errorf("Expected %d fuel but got %d instead", testCase.expected, fuel))
			}
		})
	}

	for _, testCase := range []struct {
		mass     int
		expected int
	}{
		{mass: 14, expected: 2},
		{mass: 1969, expected: 966},
		{mass: 100756, expected: 50346},
	} {
		t.Run(fmt.Sprintf("requiredTotalFuel with mass %d", testCase.mass), func(t *testing.T) {
			fuel := requiredTotalFuel(testCase.mass)
			if fuel != testCase.expected {
				t.Error(fmt.Errorf("Expected %d fuel but got %d instead", testCase.expected, fuel))
			}
		})
	}
}
