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

	output, err := exec(program, 1)
	if err != nil {
		log.Fatalf("Failed run program: %v", err)
	}

	if len(output) == 0 {
		fmt.Println("Program finished without output")
	} else {
		fmt.Printf("Program finished with output: %d", output[len(output)-1])
		fmt.Println()
	}
}

func exec(program string, input int) (Output, error) {
	memory, err := instructions(program)
	if err != nil {
		return nil, fmt.Errorf("unable to exec program: %v", err)
	}

	_, output, err := run(memory, input)
	if err != nil {
		return nil, fmt.Errorf("unable to run program: %v\nmemory: %v", err, memory)
	}
	return output, nil
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

func run(memory Memory, input int) (Memory, Output, error) {
	output := make(Output, 0)
	pc := 0
	for {
		opcode, mode1, mode2 := parseOpcode(memory.get(pc, modePosition))
		switch opcode {
		case opAdd:
			param1 := memory.get(pc+1, modePosition)
			param2 := memory.get(pc+2, modePosition)
			param3 := memory.get(pc+3, modePosition)
			result := memory.get(param1, mode1) + memory.get(param2, mode2)
			memory.set(param3, result)
			pc += 4
		case opMultiply:
			param1 := memory.get(pc+1, modePosition)
			param2 := memory.get(pc+2, modePosition)
			param3 := memory.get(pc+3, modePosition)
			result := memory.get(param1, mode1) * memory.get(param2, mode2)
			memory.set(param3, result)
			pc += 4
		case opInput:
			address := memory.get(pc+1, modePosition)
			memory.set(address, input)
			pc += 2
		case opOutput:
			address := memory.get(pc+1, modePosition)
			value := memory.get(address, mode1)
			output = append(output, value)
			pc += 2
		case opHalt:
			return memory, output, nil
		default:
			// unknown opcode
			return memory, output, fmt.Errorf("unknown opcode %d at instruction %d", opcode, pc)
		}
	}
}

// Memory represents the state of the computer
type Memory []int

func (m Memory) get(address int, mode ParameterMode) int {
	switch mode {
	case modeImmediate:
		return address
	case modePosition:
		fallthrough
	default:
		return m[address]
	}
}

func (m Memory) set(address int, value int) {
	m[address] = value
}

// Opcode represents an operation code for the intepreter
type Opcode int

const (
	opAdd      Opcode = 1
	opMultiply        = 2
	opInput           = 3
	opOutput          = 4
	opHalt            = 99
)

func parseOpcode(value int) (Opcode, ParameterMode, ParameterMode) {
	opcode := Opcode(value % 100)
	mode1 := parseMode(value / 100 & 1)
	mode2 := parseMode(value / 1000 & 1)
	return opcode, mode1, mode2
}

// Output represents the computer input
type Output []int

// ParameterMode represents the mode that defines how to interpret opcode parameters
type ParameterMode int

const (
	modePosition  ParameterMode = 0
	modeImmediate               = 1
)

func parseMode(value int) ParameterMode {
	switch value {
	case 1:
		return modeImmediate
	case 0:
		fallthrough
	default:
		return modePosition
	}
}
