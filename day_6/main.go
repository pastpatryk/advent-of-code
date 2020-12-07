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

	forms, err := readForms(file)
	check(err)

	// findEmptySeat(seats)
	// id := 8*row + col
	var sum int
	for _, form := range forms {
		sum += len(form)
	}
	fmt.Printf("Forms len: %d\n", sum)
}

func readForms(reader io.Reader) ([]map[string]bool, error) {
	var forms []map[string]bool

	form := map[string]int{}
	var lines int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			forms = append(forms, allYesForm(form, lines))
			lines = 0
			form = map[string]int{}
		} else {
			lines++
			for _, letter := range text {
				form[string(letter)]++
			}
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	forms = append(forms, allYesForm(form, lines))
	return forms, nil
}

func allYesForm(form map[string]int, count int) map[string]bool {
	yesForm := map[string]bool{}

	for key, num := range form {
		if num == count {
			yesForm[key] = true
		}
	}
	return yesForm
}
