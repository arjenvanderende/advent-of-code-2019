package main

import (
	"fmt"
	"io"
	"log"
	"os"
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

	totalMass := 0
	for {
		var mass int
		_, err = fmt.Fscanf(file, "%d", &mass)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read line from %s: %v", filename, err)
		}

		totalMass += requiredFuel(mass)
	}
	fmt.Printf("Total mass: %d", totalMass)
	fmt.Println()
}

func requiredFuel(mass int) int {
	return (mass / 3) - 2
}
