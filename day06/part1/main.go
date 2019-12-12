package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	input := string(b)
	result, err := orbits(input)
	if err != nil {
		log.Fatalf("Failed to calculate the number of orbits: %v", err)
	}
	fmt.Printf("Number of orbits: %d", result)
	fmt.Println()
}

type graph map[string][]string

func orbits(input string) (int, error) {
	g := make(graph)

	// setup orbit graph
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		objects := strings.SplitN(line, ")", 2)
		if len(objects) != 2 {
			return 0, fmt.Errorf("invalid line %s", line)
		}

		object1 := objects[0]
		object2 := objects[1]
		if _, ok := g[object1]; !ok {
			g[object1] = []string{object2}
		} else {
			g[object1] = append(g[object1], object2)
		}
	}

	return traverse(g, "COM", 0), nil
}

func traverse(g graph, node string, distance int) int {
	if _, ok := g[node]; !ok {
		return 0
	}

	sum := 0
	for _, subnode := range g[node] {
		sum += distance + 1
		sum += traverse(g, subnode, distance+1)
	}
	return sum
}
