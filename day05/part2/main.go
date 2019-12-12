package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	filename = "input.txt"
	input    = 5
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

	output, err := exec(program, input)
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
		opcode, mode1, mode2 := parseOpcode(memory.get(pc, modeImmediate))
		switch opcode {
		case opAdd:
			input1 := memory.get(pc+1, mode1)
			input2 := memory.get(pc+2, mode2)
			destination := memory.get(pc+3, modeImmediate)
			memory.set(destination, input1+input2)
			pc += 4
		case opMultiply:
			input1 := memory.get(pc+1, mode1)
			input2 := memory.get(pc+2, mode2)
			destination := memory.get(pc+3, modeImmediate)
			memory.set(destination, input1*input2)
			pc += 4
		case opInput:
			address := memory.get(pc+1, modeImmediate)
			memory.set(address, input)
			pc += 2
		case opOutput:
			value := memory.get(pc+1, mode1)
			output = append(output, value)
			pc += 2
		case opJumpIfTrue:
			value := memory.get(pc+1, mode1)
			if value != 0 {
				pc = memory.get(pc+2, mode2)
			} else {
				pc += 3
			}
		case opJumpIfFalse:
			value := memory.get(pc+1, mode1)
			if value == 0 {
				pc = memory.get(pc+2, mode2)
			} else {
				pc += 3
			}
		case opLessThan:
			value1 := memory.get(pc+1, mode1)
			value2 := memory.get(pc+2, mode2)
			destination := memory.get(pc+3, modeImmediate)
			if value1 < value2 {
				memory.set(destination, 1)
			} else {
				memory.set(destination, 0)
			}
			pc += 4
		case opEquals:
			value1 := memory.get(pc+1, mode1)
			value2 := memory.get(pc+2, mode2)
			destination := memory.get(pc+3, modeImmediate)
			if value1 == value2 {
				memory.set(destination, 1)
			} else {
				memory.set(destination, 0)
			}
			pc += 4
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
	if mode == modePosition {
		address = m[address]
	}

	return m[address]
}

func (m Memory) set(address int, value int) {
	m[address] = value
}

// Opcode represents an operation code for the intepreter
type Opcode int

const (
	opAdd         Opcode = 1
	opMultiply           = 2
	opInput              = 3
	opOutput             = 4
	opJumpIfTrue         = 5
	opJumpIfFalse        = 6
	opLessThan           = 7
	opEquals             = 8
	opHalt               = 99
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
