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

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	startTime, buses, err := readBuses(file)
	check(err)

	matchingBuses := make(map[int]int)
	for _, busID := range buses {
		if busID <= 0 {
			continue
		}
		var i int
		for i = 1; i*busID < startTime; i++ {
		}
		matchingBuses[busID] = i * busID
	}

	var bestBus int
	bestTime := math.MaxInt32
	for busID, time := range matchingBuses {
		if time < bestTime {
			bestTime = time
			bestBus = busID
		}
	}

	fmt.Printf("Best bus: %d on time %d, diff: %d\n", bestBus, bestTime, bestTime-startTime)
	fmt.Printf("Result 1: %d\n", bestBus*(bestTime-startTime))

	var multiples [][2]int
	for i, busID := range buses {
		if busID <= 0 {
			continue
		}
		multiples = append(multiples, [2]int{busID, -i})
	}

	fmt.Printf("-> %v\n", multiples)

	result := findSolution(multiples)

	fmt.Printf("Result: %d\n", result)
}

func findSolution(equations [][2]int) int {
	if len(equations) == 1 {
		return equations[0][1]
	}
	var solutions [][2]int
	for i := 1; i < len(equations); i++ {
		a, b := findLeastMultipleWithShift(equations[0][0], equations[i][0], equations[0][1], equations[i][1])
		fmt.Printf("Found: %v (%dN + %d)\n", equations[0], a, b)
		solutions = append(solutions, [2]int{a, b})
	}
	res := findSolution(solutions)
	return res*equations[0][0] + equations[0][1]
}

func readBuses(reader io.Reader) (int, []int, error) {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	time, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, err
	}
	fmt.Printf("Time: %d\n", time)

	var buses []int
	scanner.Scan()
	list := strings.Split(scanner.Text(), ",")
	for i, bus := range list {
		if bus != "x" {
			busID, err := strconv.Atoi(bus)
			if err != nil {
				return 0, nil, err
			}
			fmt.Printf("Bus: %d + %d\n", busID, i)
			buses = append(buses, busID)
		} else {
			buses = append(buses, 0)
		}
	}

	if scanner.Err() != nil {
		return 0, nil, scanner.Err()
	}
	return time, buses, nil
}

func findLeastMultipleWithShift(a, b, c, d int) (int, int) {
	e := d - c
	var i int
	for i = 1; (e+(i*b))%a != 0; i++ {
	}
	div := (e + (i * b)) / a
	fmt.Printf("[] %d / %d = %d (+ %d) = %d\n", (e + (i * b)), a, div, b, div+b)
	fmt.Printf("%dN + %d\n", b, div)
	return b, div
}
