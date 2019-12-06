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

	total := 0
	for {
		var mass int
		_, err = fmt.Fscanf(file, "%d", &mass)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to read line from %s: %v", filename, err)
		}

		total += requiredTotalFuel(mass)
	}
	fmt.Printf("Total mass: %d", total)
	fmt.Println()
}

func requiredTotalFuel(mass int) int {
	total := 0
	for {
		fuel := requiredFuel(mass)
		if fuel < 0 {
			break
		}

		total += fuel
		mass = fuel
	}
	return total
}

func requiredFuel(mass int) int {
	return (mass / 3) - 2
}
