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
	cubes  [][][][]bool
	offset [4]int
}

// Print prints visual representation
func (s *Space) Print() {
	for w, cube := range s.cubes {
		for z, level := range cube {
			fmt.Printf("\nw=%d, z=%d\n", w-s.offset[0], z-s.offset[1])
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
}

// Get returns single cube state
func (s *Space) Get(coord [4]int) bool {
	w, z, y, x := s.translateCoord(coord)
	if s.outOfBounds(w, z, y, x) {
		return false
	}
	return s.cubes[w][z][y][x]
}

// Set sets single cube state
func (s *Space) Set(coord [4]int, active bool) {
	w, z, y, x := s.translateCoord(coord)

	if !active && s.outOfBounds(w, z, y, x) {
		return
	}

	w, z, y, x = s.ensureSpaceSize(w, z, y, x)

	s.cubes[w][z][y][x] = active
}

func (s *Space) outOfBounds(w, z, y, x int) bool {
	return w < 0 || z < 0 || y < 0 || x < 0 || w >= len(s.cubes) || z >= len(s.cubes[0]) || y >= len(s.cubes[0][0]) || x >= len(s.cubes[0][0][0])
}

func (s *Space) translateCoord(coord [4]int) (int, int, int, int) {
	w := coord[0] + s.offset[0]
	z := coord[1] + s.offset[1]
	y := coord[2] + s.offset[2]
	x := coord[3] + s.offset[3]

	return w, z, y, x
}

func (s *Space) ensureSpaceSize(w, z, y, x int) (int, int, int, int) {

	if w < 0 || w >= len(s.cubes) {

		emptyCube := make([][][]bool, len(s.cubes[0]))
		for i := 0; i < len(emptyCube); i++ {
			emptySlice := make([][]bool, len(s.cubes[0][0]))
			for j := 0; j < len(emptySlice); j++ {
				emptySlice[j] = make([]bool, len(s.cubes[0][0][0]))
			}
			emptyCube[i] = emptySlice
		}

		if w < 0 {
			s.cubes = append([][][][]bool{emptyCube}, s.cubes...)
			s.offset[0] -= w
			w = 0
		} else {
			s.cubes = append(s.cubes, emptyCube)
		}
	}

	if z < 0 || z >= len(s.cubes[0]) {
		for i := 0; i < len(s.cubes); i++ {

			emptySlice := make([][]bool, len(s.cubes[0][0]))
			for j := 0; j < len(s.cubes[0][0]); j++ {
				emptySlice[j] = make([]bool, len(s.cubes[0][0][0]))
			}

			if z < 0 {
				s.cubes[i] = append([][][]bool{emptySlice}, s.cubes[i]...)
			} else {
				s.cubes[i] = append(s.cubes[i], emptySlice)
			}
		}
		if z < 0 {
			s.offset[1] -= z
			z = 0
		}
	}

	if y < 0 || y >= len(s.cubes[0][0]) {
		for _, cube := range s.cubes {
			for i := 0; i < len(cube); i++ {
				emptyRow := make([]bool, len(s.cubes[0][0][0]))
				if y < 0 {
					cube[i] = append([][]bool{emptyRow}, cube[i]...)
				} else {
					cube[i] = append(cube[i], emptyRow)
				}
			}
		}
		if y < 0 {
			s.offset[2] -= y
			y = 0
		}
	}

	if x < 0 || x >= len(s.cubes[0][0][0]) {
		for _, cube := range s.cubes {
			for _, level := range cube {
				for i := 0; i < len(level); i++ {
					if x < 0 {
						level[i] = append([]bool{false}, level[i]...)
					} else {
						level[i] = append(level[i], false)
					}
				}
			}
		}
		if x < 0 {
			s.offset[3] -= x
			x = 0
		}
	}

	return w, z, y, x
}

func neighbourCubes(coord [4]int) [][4]int {
	var coords [][4]int
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if !(w == 0 && z == 0 && y == 0 && x == 0) {
						coords = append(coords, [4]int{coord[0] + w, coord[1] + z, coord[2] + y, coord[3] + x})
					}
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

	space := Space{cubes: [][][][]bool{[][][]bool{state}}}
	space.Print()

	for cycle := 0; cycle < 6; cycle++ {
		newStates := make(map[[4]int]bool)

		for w := -1; w <= len(space.cubes); w++ {
			for z := -1; z <= len(space.cubes[0]); z++ {
				for y := -1; y <= len(space.cubes[0][0]); y++ {
					for x := -1; x <= len(space.cubes[0][0][0]); x++ {

						coord := [4]int{w - space.offset[0], z - space.offset[1], y - space.offset[2], x - space.offset[3]}

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
		}

		for coord, active := range newStates {
			space.Set(coord, active)
		}
	}

	activeCount := 0
	for _, cube := range space.cubes {
		for _, level := range cube {
			for _, row := range level {
				for _, active := range row {
					if active {
						activeCount++
					}
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
