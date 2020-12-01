package main

import (
	"bufio"
	"fmt"
	"io"
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
	n1, n2, n3 := findSum(numbers, 2020)

	fmt.Printf("Found numbers %d + %d + %d = %d\n", n1, n2, n3, n1+n2+n3)
	fmt.Printf("%d * %d * %d = %d\n", n1, n2, n3, n1*n2*n3)
}

func readNumbers(reader io.Reader) ([]int, error) {
	var numbers []int
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func findSum(numbers []int, sum int) (int, int, int) {
	for i, n1 := range numbers {
		if n1 > sum {
			continue
		}
		for j, n2 := range numbers[i:] {
			if n1+n2 > sum {
				continue
			}
			for _, n3 := range numbers[j:] {
				if n1+n2+n3 == sum {
					return n1, n2, n3
				}
			}
		}
	}
	return 0, 0, 0
}
