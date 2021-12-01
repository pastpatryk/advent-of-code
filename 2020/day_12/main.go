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
	x        int
	y        int
	waypoint [2]int
}

// MoveWaypoint changes ship's waypoint
func (s *Ship) MoveWaypoint(dir byte, val int) {
	switch dir {
	case 'N':
		s.waypoint[1] -= val
	case 'E':
		s.waypoint[0] += val
	case 'S':
		s.waypoint[1] += val
	case 'W':
		s.waypoint[0] -= val
	}
}

// RotateWaypoint rotates waypoint by a given degree
func (s *Ship) RotateWaypoint(val int) {
	deg := float64(val) * math.Pi / 180
	x := float64(s.waypoint[0])
	y := float64(s.waypoint[1])
	s.waypoint[0] = int(math.Round(x*math.Cos(deg) - y*math.Sin(deg)))
	s.waypoint[1] = int(math.Round(x*math.Sin(deg) + y*math.Cos(deg)))
}

// Move moves ship to waypoiint given amount of times
func (s *Ship) Move(val int) {
	s.x += s.waypoint[0] * val
	s.y += s.waypoint[1] * val
}

// Instruction ...
type Instruction struct {
	op  byte
	val int
}

// Execute applies instructoni to ship
func (i *Instruction) Execute(ship *Ship) {
	fmt.Printf("Executing %c %d on %v\n", i.op, i.val, ship)
	switch i.op {
	case 'R':
		ship.RotateWaypoint(i.val)
	case 'L':
		ship.RotateWaypoint(-i.val)
	case 'F':
		ship.Move(i.val)
	default:
		ship.MoveWaypoint(i.op, i.val)
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

	ship := Ship{waypoint: [2]int{10, -1}}
	for _, inst := range instructions {
		inst.Execute(&ship)
	}
	fmt.Printf("Ship position: %v (%0.f)\n", ship, math.Abs(float64(ship.x))+math.Abs(float64(ship.y)))
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
