package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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

	numbers, err := readNumbers(file)
	check(err)

	sort.Ints(numbers)

	diffs := findJoltageDiffs(numbers)
	fmt.Printf("result: %d\n", diffs[1]*diffs[3])
}

func readNumbers(reader io.Reader) ([]int, error) {
	var numbers []int
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func findJoltageDiffs(numbers []int) map[int]int {
	diffs := map[int]int{}
	prevNum := 0
	for _, num := range numbers {
		diffs[num-prevNum]++
		prevNum = num
	}
	diffs[3]++
	return diffs
}
