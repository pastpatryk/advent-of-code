package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
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

	var invalidNumber int
	for i := 25; i < len(numbers); i++ {
		if !isSum(numbers[i], numbers[i-25:]) {
			fmt.Printf("Invalid number: %d\n", numbers[i])
			invalidNumber = numbers[i]
			break
		}
	}

	foundRange := findContiguousNumbers(invalidNumber, numbers)
	fmt.Printf("range: %v\n", foundRange)
	min, max := minMax(foundRange)
	fmt.Printf("min: %d max: %d sum %d\n", min, max, min+max)
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

func findContiguousNumbers(x int, numbers []int) []int {
	for i := range numbers {
		sum := 0
		for j := i; j < len(numbers) && sum < x; j++ {
			sum += numbers[j]
			if sum == x {
				fmt.Printf("Found sum %d %d\n", sum, x)
				return numbers[i : j+1]
			}
		}
	}
	return nil
}

func isSum(x int, numbers []int) bool {
	for i, n1 := range numbers {
		for _, n2 := range numbers[i:] {
			if n1+n2 == x {
				return true
			}
		}
	}
	return false
}

func minMax(numbers []int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}
