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

// Memory represents the state of the computer
type Memory []int

func (m Memory) get(address int) int {
	return m[address]
}

func (m Memory) set(address int, value int) {
	m[address] = value
}

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

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			output, err := exec(program, noun, verb)
			if err != nil {
				log.Fatalf("Failed run program with noun %d and verb %d: %v", noun, verb, err)
			}

			if output == 19690720 {
				answer := 100*noun + verb
				fmt.Printf("Found answer: %d", answer)
				fmt.Println("")
				return
			}
		}
	}
	fmt.Println("No answer found")
}

func exec(program string, noun int, verb int) (int, error) {
	memory, err := instructions(program)
	if err != nil {
		return 0, fmt.Errorf("unable to exec program: %v", err)
	}

	// fix the program
	memory.set(1, noun)
	memory.set(2, verb)

	output, err := run(memory)
	if err != nil {
		return 0, fmt.Errorf("unable to run program: %v", err)
	}
	return output[0], nil
}

func instructions(program string) (Memory, error) {
	p := strings.Split(program, ",")
	memory := make(Memory, len(p))
	for i := range p {
		_, err := fmt.Sscanf(p[i], "%d", &memory[i])
		if err != nil {
			return memory, fmt.Errorf("instruction %d (%s) is not an opcode", i, p[i])
		}
	}
	return memory, nil
}

func run(memory Memory) (Memory, error) {
	pc := 0
	for {
		opcode := memory.get(pc)
		switch opcode {
		case 1:
			// add
			param1 := memory.get(pc + 1)
			param2 := memory.get(pc + 2)
			param3 := memory.get(pc + 3)
			result := memory.get(param1) + memory.get(param2)
			memory.set(param3, result)
			pc += 4
		case 2:
			// multiply
			param1 := memory.get(pc + 1)
			param2 := memory.get(pc + 2)
			param3 := memory.get(pc + 3)
			result := memory.get(param1) * memory.get(param2)
			memory.set(param3, result)
			pc += 4
		case 99:
			// halt
			return memory, nil
		default:
			// unknown opcode
			return memory, fmt.Errorf("unknown opcode %d at instruction %d", opcode, pc)
		}
	}
}
