package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Instruction ...
type Instruction struct {
	op   string
	arg  int
	val  int
	mask Mask
}

// Mask ...
type Mask struct {
	zeroes int
	ones   int
}

// Apply ...
func (m *Mask) Apply(number int) int {
	number |= m.ones
	number &= ^m.zeroes
	return number
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	instructions, err := readInstructions(file)
	check(err)

	fmt.Printf("Instructions: %d\n", len(instructions))

	memory := make(map[int]int)
	executeProgram(instructions, memory)

	sum := 0
	for _, val := range memory {
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}

func executeProgram(instructions []Instruction, memory map[int]int) {
	currentMask := Mask{}
	for _, inst := range instructions {
		switch inst.op {
		case "mask":
			currentMask = inst.mask
		case "mem":
			fmt.Printf("Writing to mem[%d] = %d (%d)\n", inst.arg, currentMask.Apply(inst.val), inst.val)
			memory[inst.arg] = currentMask.Apply(inst.val)
		}
	}
}

func readInstructions(reader io.Reader) ([]Instruction, error) {
	var instructions []Instruction
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return instructions, nil
}

func parseInstruction(text string) (Instruction, error) {
	inst := Instruction{}
	var err error

	sides := strings.Split(text, " = ")

	if strings.HasPrefix(sides[0], "mem") {
		inst.op = "mem"
		address := strings.TrimSuffix(strings.TrimPrefix(sides[0], "mem["), "]")
		inst.arg, err = strconv.Atoi(address)
		if err != nil {
			return inst, err
		}
		inst.val, err = strconv.Atoi(sides[1])
		if err != nil {
			return inst, err
		}
	} else {
		inst.op = "mask"
		inst.mask = parseMask(sides[1])
	}
	return inst, nil
}

func parseMask(text string) Mask {
	mask := Mask{}
	for _, char := range text {
		switch char {
		case '1':
			mask.ones |= 1
		case '0':
			mask.zeroes |= 1
		}
		mask.ones <<= 1
		mask.zeroes <<= 1
	}
	mask.ones >>= 1
	mask.zeroes >>= 1
	return mask
}
