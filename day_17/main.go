package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Space is 3dim cubes space
type Space struct {
	cubes  [][][]bool
	offset [3]int
}

// Print prints visual representation
func (s *Space) Print() {
	for z, level := range s.cubes {
		fmt.Printf("\nz=%d\n", z-s.offset[0])
		for _, row := range level {
			for _, active := range row {
				if active {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}
	}
}

// Get returns single cube state
func (s *Space) Get(coord [3]int) bool {
	z, y, x := s.translateCoord(coord)
	if s.outOfBounds(z, y, x) {
		return false
	}
	return s.cubes[z][y][x]
}

// Set sets single cube state
func (s *Space) Set(coord [3]int, active bool) {
	z, y, x := s.translateCoord(coord)

	if !active && s.outOfBounds(z, y, x) {
		return
	}

	z, y, x = s.ensureSpaceSize(z, y, x)
	s.cubes[z][y][x] = active
}

func (s *Space) outOfBounds(z, y, x int) bool {
	return z < 0 || y < 0 || x < 0 || z >= len(s.cubes) || y >= len(s.cubes[0]) || x >= len(s.cubes[0][0])
}

func (s *Space) translateCoord(coord [3]int) (int, int, int) {
	z := coord[0] + s.offset[0]
	y := coord[1] + s.offset[1]
	x := coord[2] + s.offset[2]

	return z, y, x
}

func (s *Space) ensureSpaceSize(z, y, x int) (int, int, int) {
	if z < 0 || z >= len(s.cubes) {
		emptySlice := make([][]bool, len(s.cubes[0]))
		for i := 0; i < len(s.cubes[0]); i++ {
			emptySlice[i] = make([]bool, len(s.cubes[0][0]))
		}
		if z < 0 {
			s.cubes = append([][][]bool{emptySlice}, s.cubes...)
			s.offset[0] -= z
			z = 0
		} else {
			s.cubes = append(s.cubes, emptySlice)
		}
	}

	if y < 0 || y >= len(s.cubes[0]) {
		for i := 0; i < len(s.cubes); i++ {
			emptyRow := make([]bool, len(s.cubes[0][0]))
			if y < 0 {
				s.cubes[i] = append([][]bool{emptyRow}, s.cubes[i]...)
			} else {
				s.cubes[i] = append(s.cubes[i], emptyRow)
			}
		}
		if y < 0 {
			s.offset[1] -= y
			y = 0
		}
	}

	if x < 0 || x >= len(s.cubes[0][0]) {
		for i := 0; i < len(s.cubes); i++ {
			for j := 0; j < len(s.cubes[i]); j++ {
				if x < 0 {
					s.cubes[i][j] = append([]bool{false}, s.cubes[i][j]...)
				} else {
					s.cubes[i][j] = append(s.cubes[i][j], false)
				}
			}
		}
		if x < 0 {
			s.offset[2] -= x
			x = 0
		}
	}

	return z, y, x
}

func neighbourCubes(coord [3]int) [][3]int {
	var coords [][3]int
	for z := -1; z <= 1; z++ {
		for y := -1; y <= 1; y++ {
			for x := -1; x <= 1; x++ {
				if !(z == 0 && y == 0 && x == 0) {
					coords = append(coords, [3]int{coord[0] + z, coord[1] + y, coord[2] + x})
				}
			}
		}
	}
	return coords
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	state, err := readInitState(file)
	check(err)

	space := Space{cubes: [][][]bool{state}}
	space.Print()

	for cycle := 0; cycle < 6; cycle++ {
		newStates := make(map[[3]int]bool)

		for z := -1; z <= len(space.cubes); z++ {
			for y := -1; y <= len(space.cubes[0]); y++ {
				for x := -1; x <= len(space.cubes[0][0]); x++ {
					coord := [3]int{z - space.offset[0], y - space.offset[1], x - space.offset[2]}
					activeNeighbours := 0
					for _, neighbour := range neighbourCubes(coord) {
						if space.Get(neighbour) {
							activeNeighbours++
						}
					}
					active := space.Get(coord)
					if active {
						newStates[coord] = activeNeighbours == 2 || activeNeighbours == 3
					} else {
						newStates[coord] = activeNeighbours == 3
					}
				}
			}
		}

		for coord, active := range newStates {
			space.Set(coord, active)
		}
		// fmt.Printf("=========== CYCLE: %d =============\n", cycle+1)
		// space.Print()
	}

	activeCount := 0
	for _, level := range space.cubes {
		for _, row := range level {
			for _, active := range row {
				if active {
					activeCount++
				}
			}
		}
	}

	fmt.Printf("Active count: %d\n", activeCount)
}

func readInitState(reader io.Reader) ([][]bool, error) {
	var state [][]bool
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		var row []bool
		for _, char := range line {
			row = append(row, char == '#')
		}
		state = append(state, row)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return state, nil
}
