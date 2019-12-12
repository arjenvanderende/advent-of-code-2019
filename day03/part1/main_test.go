package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

	for _, testCase := range []struct {
		line1    string
		line2    string
		expected int
	}{
		{line1: "R75,D30,R83,U83,L12,D49,R71,U7,L72", line2: "U62,R66,U55,R34,D71,R55,D58,R83", expected: 159},
		{line1: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", line2: "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", expected: 135},
	} {
		t.Run(fmt.Sprintf("closestIntersection %d", testCase.expected), func(t *testing.T) {
			actual, err := closestIntersection(testCase.line1, testCase.line2)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if testCase.expected != actual {
				t.Errorf("expected distance %d but got %d", testCase.expected, actual)
			}
		})
	}

	for _, testCase := range []struct {
		input    string
		expected move
	}{
		{input: "R75", expected: move{direction: right, distance: 75}},
		{input: "D30", expected: move{direction: down, distance: 30}},
		{input: "U83", expected: move{direction: up, distance: 83}},
		{input: "L12", expected: move{direction: left, distance: 12}},
	} {
		t.Run(fmt.Sprintf("parseMove %s", testCase.input), func(t *testing.T) {
			actual, err := parseMove(testCase.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			assertMoveEquals(t, testCase.expected, actual)
		})
	}

	for _, testCase := range []struct {
		x        int
		y        int
		expected int
	}{
		{x: 0, y: 0, expected: 0},
		{x: 1, y: 0, expected: 1},
		{x: 0, y: 1, expected: 1},
		{x: 1, y: 1, expected: 2},
		{x: -1, y: 1, expected: 2},
		{x: 1, y: -1, expected: 2},
		{x: -1, y: -1, expected: 2},
	} {
		t.Run(fmt.Sprintf("manhattanDistance %d, %d", testCase.x, testCase.y), func(t *testing.T) {
			actual := manhattanDistance(coord{x: testCase.x, y: testCase.y})
			if testCase.expected != actual {
				t.Errorf("expected %d but got %d", testCase.expected, actual)
			}
		})
	}
}

func assertMoveEquals(t *testing.T, expected move, actual move) {
	if expected.direction != actual.direction {
		t.Errorf("expected direction: %v but got: %v", expected.direction, actual.direction)
	}
	if expected.distance != actual.distance {
		t.Errorf("expected distance: %d but got: %d", expected.distance, actual.distance)
	}
}
