package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type passwordPolicy struct {
	letter string
	min    int
	max    int
}

func (p passwordPolicy) isValid(password string) bool {
	count := strings.Count(password, p.letter)
	return p.min <= count && count <= p.max
}

type passwordDefinition struct {
	policy   passwordPolicy
	password string
}

func (pd passwordDefinition) isValid() bool {
	return pd.policy.isValid(pd.password)
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	passwords, err := readPasswords(file)
	check(err)
	fmt.Printf("Number of passwords: %d\n", len(passwords))

	var validPasswords []passwordDefinition
	for _, password := range passwords {
		if password.isValid() {
			validPasswords = append(validPasswords, password)
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", len(validPasswords))
}

func readPasswords(reader io.Reader) ([]passwordDefinition, error) {
	var passwords []passwordDefinition
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		definition, err := parseDefinition(line)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, definition)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return passwords, nil
}

func parseDefinition(text string) (passwordDefinition, error) {
	var err error
	var passDef passwordDefinition

	definitions := strings.Split(text, " ")
	if len(definitions) < 3 {
		return passDef, fmt.Errorf("invalid input: %s", text)
	}

	occurences := strings.Split(definitions[0], "-")
	passDef.policy.min, err = strconv.Atoi(occurences[0])
	if err != nil {
		return passDef, err
	}
	passDef.policy.max, err = strconv.Atoi(occurences[1])
	if err != nil {
		return passDef, err
	}

	passDef.policy.letter = strings.TrimSuffix(definitions[1], ":")
	passDef.password = definitions[2]
	return passDef, nil
}
