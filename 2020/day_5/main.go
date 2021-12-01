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

	findEmptySeat(seats)
	// id := 8*row + col
	// fmt.Printf("Seat %d: [%d, %d]\n", max, row, col)
}

func readSeats(reader io.Reader) ([]string, error) {
	var seats []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return seats, nil
}

func decodeSeatPosition(seat string) (int, int) {
	row, _ := binarySearch(seat[:7], 0, 127)
	col, _ := binarySearch(seat[7:], 0, 7)
	return row, col
}

func binarySearch(code string, from, to int) (int, int) {
	if len(code) == 1 {
		if code[0] == 'F' || code[0] == 'L' {
			return from, from
		}
		return to, to
	}
	mid := (from + to) / 2
	if code[0] == 'F' || code[0] == 'L' {
		return binarySearch(code[1:], from, mid)
	}
	return binarySearch(code[1:], mid+1, to)
}

func findEmptySeat(seats []string) {
	max := 0
	var occupied [128][8]bool
	for _, seat := range seats {
		row, col := decodeSeatPosition(seat)
		// fmt.Printf("seat: %d, %d\n", row, col)
		occupied[row][col] = true
		id := row*8 + col
		if id > max {
			max = id
		}
	}
	var freeSeatsIds []int
	for r := range occupied {
		for c := range occupied[r] {
			if !occupied[r][c] {
				freeSeatsIds = append(freeSeatsIds, r*8+c)
				fmt.Printf("%d\n", r*8+c)
			}
		}
	}
}
