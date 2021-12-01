package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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
	zeroes      int
	ones        int
	floatingPos []int
}

// Apply ...
func (m *Mask) Apply(number int) int {
	number |= m.ones
	// number &= ^m.zeroes
	return number
}

// MemoryAdresses generates all possible memeory adresses
func (m *Mask) MemoryAdresses(address int) []int {
	combinations := powerSet(m.floatingPos)
	var addresses []int
	for _, combination := range combinations {
		newAddress := address
		for pos, flag := range combination {
			if flag {
				newAddress = setBit(newAddress, uint(pos))
			} else {
				newAddress = clearBit(newAddress, uint(pos))
			}
		}
		addresses = append(addresses, m.Apply(newAddress))
	}
	return addresses
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func clearBit(n int, pos uint) int {
	return n &^ (1 << pos)
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	fmt.Printf("Powerset: %v\n", powerSet([]int{1, 2}))

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
	for i, inst := range instructions {
		switch inst.op {
		case "mask":
			currentMask = inst.mask
		case "mem":
			addresses := currentMask.MemoryAdresses(inst.arg)
			for _, address := range addresses {
				fmt.Printf("[%d] Writing to mem[%d](%d) = %d\n", i, address, inst.arg, inst.val)
				memory[address] = inst.val
			}
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
	for i, char := range text {
		switch char {
		case '1':
			mask.ones |= 1
		case '0':
			mask.zeroes |= 1
		case 'X':
			mask.floatingPos = append(mask.floatingPos, len(text)-i-1)
		}
		mask.ones <<= 1
		mask.zeroes <<= 1
	}
	mask.ones >>= 1
	mask.zeroes >>= 1
	return mask
}

func powerSet(original []int) []map[int]bool {
	powerSetSize := int(math.Pow(2, float64(len(original))))
	result := make([]map[int]bool, 0, powerSetSize)

	var index int
	for index < powerSetSize {
		subSet := make(map[int]bool, len(original))

		for j, elem := range original {
			subSet[elem] = index&(1<<uint(j)) > 0
		}
		result = append(result, subSet)
		index++
	}
	return result
}
