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

type BagCount struct {
	color string
	count int
}

type BagRule struct {
	color string
	bags  []BagCount
}

func (b *BagRule) directlyContainsColor(color string) bool {
	for _, bag := range b.bags {
		if bag.color == color {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	rules, err := readRules(file)
	check(err)

	count := countBagContaining(rules, "shiny gold")
	required := countRequiredBags(rules, "shiny gold")
	fmt.Printf("Count: %d required %d\n", count, required)
	// findEmptySeat(seats)
	// id := 8*row + col
	// fmt.Printf("Seat %d: [%d, %d]\n", max, row, col)
}

func readRules(reader io.Reader) (map[string]BagRule, error) {
	rules := map[string]BagRule{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		rule, err := parseRule(scanner.Text())
		if err != nil {
			return nil, err
		}
		rules[rule.color] = rule
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return rules, nil
}

func parseRule(text string) (BagRule, error) {
	rule := BagRule{}
	text = strings.TrimSuffix(text, ".")
	slices := strings.Split(text, " bags contain ")
	rule.color = slices[0]
	if slices[1] == "no other bags" {
		return rule, nil
	}

	bagsSlices := strings.Split(slices[1], ", ")
	for _, bagSlice := range bagsSlices {
		bag, err := parseBagCount(bagSlice)
		if err != nil {
			return BagRule{}, err
		}
		rule.bags = append(rule.bags, bag)
	}
	return rule, nil
}

func parseBagCount(text string) (BagCount, error) {
	var err error
	bag := BagCount{}

	slices := strings.SplitN(text, " ", 2)
	bag.count, err = strconv.Atoi(slices[0])
	if err != nil {
		return BagCount{}, err
	}
	color := strings.TrimSuffix(slices[1], " bags")
	color = strings.TrimSuffix(color, " bag")
	bag.color = color

	return bag, nil
}

func countBagContaining(rules map[string]BagRule, color string) int {
	var count int
	for _, rule := range rules {
		if containsBag(rules, rule.color, color) {
			count++
		}
	}
	return count
}

func containsBag(rules map[string]BagRule, outer string, inner string) bool {
	rule, _ := rules[outer]
	if rule.directlyContainsColor(inner) {
		return true
	}

	for _, bag := range rule.bags {
		contains := containsBag(rules, bag.color, inner)
		if contains {
			return true
		}
	}

	return false
}

func countRequiredBags(rules map[string]BagRule, color string) int {
	rule, _ := rules[color]
	fmt.Printf("%v\n", rule)

	var required int
	for _, bag := range rule.bags {
		required += bag.count
		required += bag.count * countRequiredBags(rules, bag.color)
	}
	return required
}
