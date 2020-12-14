package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// Ship represents position and direction of a ship
type Ship struct {
	x      int
	y      int
	facing int
}

// Move changes ship's position
func (s *Ship) Move(dir byte, val int) {
	switch dir {
	case 'N':
		s.y -= val
	case 'E':
		s.x += val
	case 'S':
		s.y += val
	case 'W':
		s.x -= val
	}
}

// Instruction ...
type Instruction struct {
	op  byte
	val int
}

// Execute applies instructoni to ship
func (i *Instruction) Execute(ship *Ship) {
	fmt.Printf("Executing %c %d on %v\n", i.op, i.val, ship)
	dirs := []byte{'N', 'E', 'S', 'W'}
	switch i.op {
	case 'R':
		ship.facing = (ship.facing + (i.val / 90) + len(dirs)) % len(dirs)
	case 'L':
		ship.facing = (ship.facing - (i.val / 90) + len(dirs)) % len(dirs)
	case 'F':
		ship.Move(dirs[ship.facing], i.val)
	default:
		ship.Move(i.op, i.val)
	}
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

	instructions, err := readInstructions(file)
	check(err)

	ship := Ship{facing: 1}
	for _, inst := range instructions {
		inst.Execute(&ship)
	}
	fmt.Printf("Ship position: %v (%f)\n", ship, math.Abs(float64(ship.x))+math.Abs(float64(ship.y)))
}

func readInstructions(reader io.Reader) ([]Instruction, error) {
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
	instruction.op = s[0]

	val, err := strconv.Atoi(s[1:])
	if err != nil {
		return Instruction{}, err
	}
	instruction.val = val

	return instruction, nil
}
