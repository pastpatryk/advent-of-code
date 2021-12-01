package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Food represents single food with its ingredients
type Food struct {
	ingredients []string
	allergens   []string
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	foods, err := readFoods(file)

	fmt.Printf("Foods: %d\n", len(foods))
	fmt.Printf("Foods: %v\n", foods[0])

	ingredients := make(map[string]int)
	alleToIngr := make(map[string][]Food)
	for _, food := range foods {
		for _, alle := range food.allergens {
			alleToIngr[alle] = append(alleToIngr[alle], food)
		}
		for _, ingr := range food.ingredients {
			ingredients[ingr]++
		}
	}

	matchingIngr := make(map[string][]string)
	for alle, foods := range alleToIngr {
		ingrCounts := make(map[string]int)
		for _, food := range foods {
			for _, ingr := range food.ingredients {
				ingrCounts[ingr]++
			}
		}

		for ingr, count := range ingrCounts {
			if count == len(foods) {
				matchingIngr[alle] = append(matchingIngr[alle], ingr)
			}
		}
		fmt.Printf("Alle: %s, matching ingr: %v\n", alle, matchingIngr[alle])
	}

	final := make(map[string]string)

	for len(matchingIngr) > 0 {
		for alle, ingrs := range matchingIngr {
			if len(ingrs) == 1 {
				fmt.Printf("Alle: %q, ingr: %q\n", alle, ingrs[0])
				final[alle] = ingrs[0]
				delete(matchingIngr, alle)
				delete(ingredients, ingrs[0])
				clearFromPossibleIngredients(matchingIngr, ingrs[0])
			}
		}
	}

	sum := 0
	for _, count := range ingredients {
		sum += count
	}
	fmt.Printf("non alergic sum: %d\n", sum)

	var allergens []string
	for alle := range final {
		allergens = append(allergens, alle)
	}
	sort.Strings(allergens)

	var sortedIngredients []string
	for _, alle := range allergens {
		sortedIngredients = append(sortedIngredients, final[alle])
	}

	fmt.Printf("Result: %s\n", strings.Join(sortedIngredients, ","))
}

func clearFromPossibleIngredients(possible map[string][]string, ingr string) {
	for alle, ingrs := range possible {
		idx := findIndex(ingrs, ingr)
		if idx >= 0 {
			possible[alle] = append(ingrs[:idx], ingrs[(idx+1):]...)
		}
	}
}

func findIndex(list []string, item string) int {
	for i, el := range list {
		if el == item {
			return i
		}
	}
	return -1
}

func readFoods(reader io.Reader) ([]Food, error) {
	var foods []Food

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		foods = append(foods, parseFood(scanner.Text()))
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return foods, nil
}

func parseFood(s string) Food {
	food := Food{}
	sections := strings.Split(s, " (contains ")
	food.ingredients = strings.Split(sections[0], " ")
	food.allergens = strings.Split(strings.TrimSuffix(sections[1], ")"), ", ")

	return food
}
