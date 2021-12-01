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

// TicketField represents single field definition with it's rules
type TicketField struct {
	name   string
	ranges [][2]int
}

// Ticket is a set of field values
type Ticket []int

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	tickets, fields, err := readTickets(file)
	check(err)

	sum := 0
	var validTickets []Ticket
	for _, ticket := range tickets[1:] {
		if valid, num := validateTicket(ticket, fields); !valid {
			sum += num
		} else {
			validTickets = append(validTickets, ticket)
		}
	}

	fmt.Printf("Tickets: %d, fields: %v, valid: %d\n", len(tickets), len(fields), len(validTickets))
	fmt.Printf("Sum invalid: %d\n", sum)

	fieldPositions := make(map[string][]int)

	for _, field := range fields {
		fieldPositions[field.name] = make([]int, 0, 3)
		for i := 0; i < len(tickets[0]); i++ {
			if validateTicketsColumn(i, validTickets, field) {
				fieldPositions[field.name] = append(fieldPositions[field.name], i)
			}
		}
	}

	removed := true
	for removed {
		removed = false
		for name, positions := range fieldPositions {
			if len(positions) == 1 {
				pos := positions[0]
				fmt.Printf("Single pos %s: %d\n", name, pos)
				rem := clearPositionFromOtherField(fieldPositions, pos, name)
				if rem && !removed {
					removed = true
				}
			}
		}
	}

	var departurePositions []int
	for name, positions := range fieldPositions {
		if strings.Contains(name, "departure") {
			departurePositions = append(departurePositions, positions[0])
		}
	}

	fmt.Printf("Positions: %v\n", fieldPositions)

	myTicket := tickets[0]
	mul := 1
	for _, depPos := range departurePositions {
		mul *= myTicket[depPos]
	}

	fmt.Printf("Result: %d\n", mul)
}

func clearPositionFromOtherField(fieldPositions map[string][]int, pos int, name string) bool {
	var removed bool
	for n, positions := range fieldPositions {
		idx := find(positions, pos)
		if n != name && idx >= 0 {
			removed = true
			fmt.Printf("Removing %d from %s %v\n", pos, n, positions)
			fieldPositions[n] = append(positions[:idx], positions[(idx+1):]...)
		}
	}
	return removed
}

func find(slice []int, num int) int {
	for idx, n := range slice {
		if n == num {
			return idx
		}
	}
	return -1
}

func validateTicketsColumn(col int, tickets []Ticket, field TicketField) bool {
	for _, ticket := range tickets {
		if !validateRanges(ticket[col], field.ranges) {
			return false
		}
	}
	return true
}

func validateTicket(ticket Ticket, fields []TicketField) (bool, int) {
	for _, num := range []int(ticket) {
		if !validateFields(num, fields) {
			return false, num
		}
	}
	return true, 0
}

func validateFields(num int, fields []TicketField) bool {
	for _, field := range fields {
		if validateRanges(num, field.ranges) {
			return true
		}
	}
	return false
}

func validateRanges(num int, ranges [][2]int) bool {
	for _, r := range ranges {
		if num >= r[0] && num <= r[1] {
			return true
		}
	}
	return false
}

func readTickets(reader io.Reader) ([]Ticket, []TicketField, error) {
	var fields []TicketField
	var tickets []Ticket

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		field, err := parseField(text)
		if err != nil {
			return nil, nil, err
		}
		fields = append(fields, field)
	}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" || strings.Contains(text, "ticket") {
			continue
		}
		ticket, err := parseTicket(text)
		if err != nil {
			return nil, nil, err
		}
		tickets = append(tickets, ticket)
	}

	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}
	return tickets, fields, nil
}

func parseField(s string) (TicketField, error) {
	field := TicketField{}
	sections := strings.Split(s, ": ")
	field.name = sections[0]
	ranges := strings.Split(sections[1], " or ")
	for _, rangeString := range ranges {
		r, err := parseRange(rangeString)
		if err != nil {
			return field, err
		}
		field.ranges = append(field.ranges, r)
	}

	return field, nil
}

func parseRange(s string) ([2]int, error) {
	var fieldRange [2]int
	var err error
	numbers := strings.Split(s, "-")
	fieldRange[0], err = strconv.Atoi(numbers[0])
	if err != nil {
		return fieldRange, err
	}
	fieldRange[1], err = strconv.Atoi(numbers[1])
	if err != nil {
		return fieldRange, err
	}
	return fieldRange, nil
}

func parseTicket(s string) (Ticket, error) {
	var ticket []int
	for _, numStr := range strings.Split(s, ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return Ticket{}, err
		}
		ticket = append(ticket, num)
	}
	return Ticket(ticket), nil
}
