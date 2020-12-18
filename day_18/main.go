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

// Expression ...
type Expression []Value

// Value ...
type Value struct {
	num           int
	subexpression Expression
	op            byte
}

func (v Value) String() string {
	if len(v.subexpression) > 0 {
		return fmt.Sprintf("(%v) %c", v.subexpression, v.op)
	}
	return fmt.Sprintf("%d %c", v.num, v.op)
}

func (v Value) Eval() int {
	if len(v.subexpression) > 0 {
		return v.subexpression.Eval()
	}
	return v.num
}

func (e Expression) String() string {
	var sb strings.Builder
	for _, val := range e {
		sb.WriteString(val.String() + " ")
	}
	return sb.String()
}

// Eval calculates result of the expression
func (e Expression) Eval() int {
	result := e[0].Eval()
	op := e[0].op
	for _, val := range e[1:] {
		switch op {
		case '+':
			result += val.Eval()
		case '*':
			result *= val.Eval()
		}
		op = val.op
	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	expressions, err := readExpressions(file)
	check(err)

	sum := 0
	for _, exp := range expressions {
		// fmt.Printf("Exp: %v = %d\n", exp, exp.Eval())
		sum += exp.Eval()
	}
	fmt.Printf("Sum: %d\n", sum)
}

func readExpressions(reader io.Reader) ([]Expression, error) {
	var expressions []Expression
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		expression, err := parseExpression(scanner.Text())
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, expression)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return expressions, nil
}

func parseExpression(line string) (Expression, error) {
	line = strings.ReplaceAll(line, "(", "( ")
	line = strings.ReplaceAll(line, ")", " )")
	fragments := strings.Split(line, " ")
	expression, _, err := parseFragments(fragments)
	if err != nil {
		return nil, err
	}
	return expression, nil
}

func parseFragments(fragments []string) (Expression, []string, error) {
	var err error
	expression := Expression{}
	value := Value{}
	for len(fragments) > 0 {
		frag := fragments[0]
		fragments = fragments[1:]
		if frag == "+" || frag == "*" {
			value.op = frag[0]
			expression = append(expression, value)
			value = Value{}
		} else if frag == "(" {
			value.subexpression, fragments, err = parseFragments(fragments)
			if err != nil {
				return expression, nil, err
			}
		} else if frag == ")" {
			expression = append(expression, value)
			return expression, fragments, nil
		} else {
			value.num, err = strconv.Atoi(frag)
			if err != nil {
				return expression, nil, err
			}
		}
	}
	expression = append(expression, value)

	return expression, nil, nil
}
