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

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	seats, err := readSeats(file)
	check(err)

	newSeats, changes := generateNextState(seats)
	iter := 0
	for changes > 0 {
		fmt.Printf("Changes: %d\n", changes)
		newSeats, changes = generateNextState(newSeats)
		iter++
	}

	total := 0
	for _, row := range newSeats {
		for _, seat := range row {
			if seat == '#' {
				total++
			}
		}
	}

	fmt.Printf("Iter: %d\n", iter)
	fmt.Printf("Total: %d\n", total)
}

func readSeats(reader io.Reader) ([][]byte, error) {
	var seats [][]byte
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		seats = append(seats, row)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return seats, nil
}

func generateNextState(seats [][]byte) ([][]byte, int) {
	var newSeats [][]byte
	changes := 0
	for i, row := range seats {
		newRow := make([]byte, len(row))
		for j, seat := range row {
			adjacent := 0
			if checkLineOfSight(seats, i, j, -1, -1) == '#' {
				adjacent++
			}
			if checkLineOfSight(seats, i, j, -1, 0) == '#' {
				adjacent++
			}
			if checkLineOfSight(seats, i, j, -1, 1) == '#' {
				adjacent++
			}
			if checkLineOfSight(seats, i, j, 0, -1) == '#' {
				adjacent++
			}
			if checkLineOfSight(seats, i, j, 0, 1) == '#' {
				adjacent++
			}
			if checkLineOfSight(seats, i, j, 1, -1) == '#' {
				adjacent++
			}
			if checkLineOfSight(seats, i, j, 1, 0) == '#' {
				adjacent++
			}
			if checkLineOfSight(seats, i, j, 1, 1) == '#' {
				adjacent++
			}

			if seat == 'L' && adjacent == 0 {
				newRow[j] = '#'
				changes++
			} else if seat == '#' && adjacent >= 5 {
				newRow[j] = 'L'
				changes++
			} else {
				newRow[j] = seat
			}
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats, changes
}

func checkLineOfSight(seats [][]byte, i, j, dirX, dirY int) byte {
	i += dirX
	j += dirY
	for i >= 0 && i < len(seats) && j >= 0 && j < len(seats[i]) {
		if seats[i][j] != '.' {
			return seats[i][j]
		}
		i += dirX
		j += dirY
	}
	return '.'
}
