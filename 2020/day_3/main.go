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

// WorldMap represents positions of trees
type WorldMap [][]byte

func (w WorldMap) isTree(x, y int) bool {
	row := [][]byte(w)[y]
	return row[x%len(row)] == '#'
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	worldMap, err := readWorldMap(file)
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	trees := 1
	for _, slope := range slopes {
		trees *= countTrees(worldMap, slope)
	}
	fmt.Printf("Trees: %d\n", trees)
}

func readWorldMap(reader io.Reader) (WorldMap, error) {
	var worldMap [][]byte
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		row := []byte(scanner.Text())
		worldMap = append(worldMap, row)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return WorldMap(worldMap), nil
}

func countTrees(worldMap WorldMap, slope [2]int) int {
	x := 0
	y := 0
	trees := 0
	for y < len(worldMap) {
		if worldMap.isTree(x, y) {
			trees++
		}
		x += slope[0]
		y += slope[1]
	}
	fmt.Printf("%v - trees: %d\n", slope, trees)
	return trees
}
