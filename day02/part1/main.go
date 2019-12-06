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

	program := string(b)
	output, err := exec(program)
	if err != nil {
		log.Fatalf("Failed run program: %v", err)
	}

	fmt.Printf("Program output: %d", output)
	fmt.Println()
}

func exec(program string) (int, error) {
	opcodes, err := instructions(program)
	if err != nil {
		return 0, fmt.Errorf("unable to exec program: %v", err)
	}

	// fix the program
	opcodes[1] = 12
	opcodes[2] = 2

	output := run(opcodes)
	return output[0], nil
}

func instructions(program string) ([]int, error) {
	p := strings.Split(program, ",")
	opcodes := make([]int, len(p))
	for i := range p {
		_, err := fmt.Sscanf(p[i], "%d", &opcodes[i])
		if err != nil {
			return opcodes, fmt.Errorf("instruction %d (%s) is not an opcode", i, p[i])
		}
	}
	return opcodes, nil
}

func run(opcodes []int) []int {
	pc := 0
	for {
		opcode := opcodes[pc]
		switch opcode {
		case 1:
			// add
			arg1 := opcodes[pc+1]
			arg2 := opcodes[pc+2]
			arg3 := opcodes[pc+3]
			opcodes[arg3] = opcodes[arg1] + opcodes[arg2]
			pc += 4
		case 2:
			// multiply
			arg1 := opcodes[pc+1]
			arg2 := opcodes[pc+2]
			arg3 := opcodes[pc+3]
			opcodes[arg3] = opcodes[arg1] * opcodes[arg2]
			pc += 4
		default:
			// halt on unknown opcode
			return opcodes
		}
	}
}
