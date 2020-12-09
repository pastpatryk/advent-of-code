package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Instruction ...
type Instruction struct {
	op    string
	value int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	program, err := readProgram(file)
	check(err)

	programs := generateProgramVersions(program)

	var acc int
	for idx, prog := range programs {
		fmt.Printf("Running program %d\n", idx)
		acc, err = run(prog)
		if err == nil {
			break
		}
	}

	check(err)
	fmt.Printf("Acc: %d\n", acc)
}

func readProgram(reader io.Reader) ([]Instruction, error) {
	var instructions []Instruction
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		inst, err := parseInstruction(scanner.Text())
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, inst)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return instructions, nil
}

func parseInstruction(s string) (Instruction, error) {
	instruction := Instruction{}
	slices := strings.Split(s, " ")
	instruction.op = slices[0]

	val, err := strconv.Atoi(strings.TrimPrefix(slices[1], "+"))
	if err != nil {
		return Instruction{}, err
	}
	instruction.value = val

	return instruction, nil
}

func run(program []Instruction) (int, error) {
	var idx, acc int

	instMem := map[int]bool{}

	for idx < len(program) {
		if instMem[idx] {
			return acc, fmt.Errorf("instruction used twice")
		}
		instMem[idx] = true

		inst := program[idx]
		switch inst.op {
		case "acc":
			acc += inst.value
		case "jmp":
			idx += inst.value
			continue
		}
		idx++
	}
	return acc, nil
}

func generateProgramVersions(program []Instruction) [][]Instruction {
	var programs [][]Instruction

	for idx, inst := range program {
		switch inst.op {
		case "nop":
			newProgram := make([]Instruction, len(program))
			copy(newProgram, program)
			newProgram[idx].op = "jmp"
			programs = append(programs, newProgram)
		case "jmp":
			newProgram := make([]Instruction, len(program))
			copy(newProgram, program)
			newProgram[idx].op = "nop"
			programs = append(programs, newProgram)
		}
	}
	return programs
}
