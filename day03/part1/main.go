package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

var (
	filename = "input.txt"
)

func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", filename, err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", filename, err)
	}

	lines := strings.SplitN(string(b), "\n", 2)
	if len(lines) < 2 {
		log.Fatalf("File %s must contain at least 2 lines", filename)
	}

	distance, err := closestIntersection(lines[0], lines[1])
	if err != nil {
		log.Fatalf("unable to get closest intersection: %v", err)
	}
	fmt.Printf("Closest intersection: %d", distance)
}

func closestIntersection(line1 string, line2 string) (int, error) {
	port := &centralPort{
		wires:         make(map[coord]int),
		intersections: make([]coord, 0),
	}

	moves1, err := parseMoves(line1)
	if err != nil {
		log.Fatalf("failed to parse moves from line 1: %v", err)
	}
	port.fill(moves1, 1)

	moves2, err := parseMoves(line2)
	if err != nil {
		log.Fatalf("failed to parse moves from line 2: %v", err)
	}
	port.fill(moves2, 2)

	return port.closestIntersection()
}

type centralPort struct {
	wires         map[coord]int
	intersections []coord
}

func (port *centralPort) fill(moves []move, identifier int) {
	start := coord{x: 0, y: 0}
	for _, move := range moves {
		end := destination(start, move)
		switch move.direction {
		case up:
			for y := start.y + 1; y <= end.y; y++ {
				port.set(coord{x: start.x, y: y}, identifier)
			}
		case down:
			for y := start.y - 1; y >= end.y; y-- {
				port.set(coord{x: start.x, y: y}, identifier)
			}
		case left:
			for x := start.x - 1; x >= end.x; x-- {
				port.set(coord{x: x, y: start.y}, identifier)
			}
		case right:
			for x := start.x + 1; x <= end.x; x++ {
				port.set(coord{x: x, y: start.y}, identifier)
			}
		}
		start = end
	}
}

func (port *centralPort) set(pos coord, identifier int) {
	if elem, ok := port.wires[pos]; !ok || elem&identifier == identifier {
		port.wires[pos] = identifier
	} else {
		port.wires[pos] += identifier
		port.intersections = append(port.intersections, pos)
	}
}

func (port *centralPort) closestIntersection() (int, error) {
	if len(port.intersections) == 0 {
		return 0, fmt.Errorf("no intersections")
	}

	closest := math.MaxInt32
	for _, intersection := range port.intersections {
		distance := manhattanDistance(intersection)
		if distance < closest {
			closest = distance
		}
	}
	return closest, nil
}

func destination(start coord, move move) coord {
	switch move.direction {
	case up:
		return coord{x: start.x, y: start.y + move.distance}
	case down:
		return coord{x: start.x, y: start.y - move.distance}
	case left:
		return coord{x: start.x - move.distance, y: start.y}
	case right:
	}
	return coord{x: start.x + move.distance, y: start.y}
}

type coord struct {
	x, y int
}

func manhattanDistance(pos coord) int {
	return abs(pos.x) + abs(pos.y)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

type move struct {
	direction direction
	distance  int
}

func parseMoves(text string) ([]move, error) {
	parts := strings.Split(text, ",")
	moves := make([]move, len(parts))
	for i, part := range parts {
		move, err := parseMove(part)
		if err != nil {
			return moves, fmt.Errorf("unable to parse move %d %s: %v", i, part, err)
		}
		moves[i] = move
	}
	return moves, nil
}

func parseMove(text string) (move, error) {
	var direction string
	var distance int
	_, err := fmt.Sscanf(text, "%1s%d", &direction, &distance)
	if err != nil {
		return move{}, fmt.Errorf("unable to parse move %s: %v", text, err)
	}

	switch direction {
	case "U":
		return move{direction: up, distance: distance}, nil
	case "D":
		return move{direction: down, distance: distance}, nil
	case "L":
		return move{direction: left, distance: distance}, nil
	case "R":
		return move{direction: right, distance: distance}, nil
	}

	return move{}, fmt.Errorf("unknown direction %s", direction)
}

type direction int

const (
	up direction = iota
	down
	left
	right
)
