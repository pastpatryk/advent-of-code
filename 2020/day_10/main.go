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
	fmt.Printf("Result: %d\n", diffs[1]*diffs[3])

	numbers = append([]int{0}, numbers...)

	var splitNumbers [][]int
	var current []int
	prevNum := 0
	for _, num := range numbers {
		if num-prevNum >= 3 {
			splitNumbers = append(splitNumbers, current)
			current = []int{}
		}

		current = append(current, num)
		prevNum = num
	}
	splitNumbers = append(splitNumbers, current)

	total := 1
	for i, split := range splitNumbers {
		var max int
		if i+1 < len(splitNumbers) {
			max = splitNumbers[i+1][0]
		} else {
			max = numbers[len(numbers)-1] + 3
		}

		combinations := generateCombinations(split[0:1], split[1:len(split)], max)
		total *= len(combinations)
	}

	fmt.Printf("Total: %d\n", total)
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

func generateCombinations(current []int, available []int, target int) [][]int {
	var combinations [][]int
	currentJolts := current[len(current)-1]
	if len(available) == 0 && currentJolts+3 >= target {
		combinations = append(combinations, current)
		return combinations
	}
	for i := 0; i < len(available) && available[i] <= currentJolts+3; i++ {
		newCombination := append(current, available[i])
		combinations = append(
			combinations,
			generateCombinations(
				append([]int(nil), newCombination...),
				append([]int(nil), available[i+1:]...),
				target,
			)...,
		)
	}
	return combinations
}
