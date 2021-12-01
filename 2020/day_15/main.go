package main

import "fmt"

func main() {
	numbers := []int{8, 13, 1, 0, 18, 9}
	memory := initMemory(numbers)

	lastNum := numbers[len(numbers)-1]
	for i := len(numbers); i < 30000000; i++ {
		age, ok := memory[lastNum]
		if !ok {
			age = 0
		} else {
			age = i - age - 1
		}
		memory[lastNum] = i - 1
		lastNum = age
	}
	fmt.Printf("Last num: %d\n", lastNum)
}

func initMemory(numbers []int) map[int]int {
	memory := make(map[int]int)
	for i := 0; i < len(numbers)-1; i++ {
		memory[numbers[i]] = i
	}
	return memory
}
