package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// GrammarRule keeps single grammar rule, either letter or reference to other rules
type GrammarRule struct {
	references [][]int
	letter     string
}

// Grammar defines whole grammar
type Grammar map[int]GrammarRule

// RuleStr returns regular expression string for grammar rule
func (g Grammar) ruleStr(id int) string {
	rule := g[id]
	if rule.letter != "" {
		return rule.letter
	}

	if id == 8 {
		return fmt.Sprintf("(%s+)", g.ruleStr(42))
	}

	if id == 11 {
		var repetitions []string
		r42 := g.ruleStr(42)
		r31 := g.ruleStr(31)
		for i := 1; i <= 50; i++ {
			repetitions = append(repetitions, fmt.Sprintf("(%s{%d}%s{%d})", r42, i, r31, i))
		}
		return fmt.Sprintf("(%s)", strings.Join(repetitions, "|"))
	}

	var regexpParts []string
	for _, option := range rule.references {
		var res string
		for _, refID := range option {
			res += g.ruleStr(refID)
		}
		regexpParts = append(regexpParts, res)
	}

	return fmt.Sprintf("(%s)", strings.Join(regexpParts, "|"))
}

// Regexp returns compiled regex for grammar
func (g Grammar) Regexp() (*regexp.Regexp, error) {
	return regexp.Compile(fmt.Sprintf("^%s$", g.ruleStr(0)))
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	grammar, messages, err := readGrammar(file)
	check(err)

	fmt.Printf("Grammar len: %d, messages: %d\n", len(grammar), len(messages))

	regExp, err := grammar.Regexp()
	check(err)

	valid := 0
	for _, message := range messages {
		if regExp.MatchString(message) {
			valid++
		}
	}

	fmt.Printf("Valid messages: %d\n", valid)
}

func readGrammar(reader io.Reader) (Grammar, []string, error) {
	grammar := make(map[int]GrammarRule)
	var messages []string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		id, rule, err := parseRule(text)
		if err != nil {
			return nil, nil, err
		}
		grammar[id] = rule
	}

	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}
	return grammar, messages, nil
}

func parseRule(s string) (int, GrammarRule, error) {
	rule := GrammarRule{}
	defintions := strings.Split(s, ": ")
	id, err := strconv.Atoi(defintions[0])
	if err != nil {
		return 0, rule, err
	}

	if strings.HasPrefix(defintions[1], "\"") {
		rule.letter = defintions[1][1:2]
	} else {
		for _, option := range strings.Split(defintions[1], " | ") {
			var ids []int
			for _, idStr := range strings.Split(option, " ") {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					return 0, rule, err
				}
				ids = append(ids, id)
			}
			rule.references = append(rule.references, ids)
		}
	}

	return id, rule, nil
}
