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

// Deck is a pile of cards
type Deck []int

// Pop removes and returns first card from the deck
func (d *Deck) Pop() int {
	if len(*d) == 0 {
		return 0
	}
	card := (*d)[0]
	*d = (*d)[1:]

	return card
}

// Add adds 2 new card at the end of the deck
func (d *Deck) Add(card1, card2 int) {
	*d = append(*d, card1, card2)
}

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	deck1, deck2, err := readDecks(file)
	check(err)

	// fmt.Printf("Player1: %v\n", deck1)
	// fmt.Printf("Player2: %v\n", deck2)

	rounds := 0
	for len(deck1) > 0 && len(deck2) > 0 {
		rounds++
		card1 := deck1.Pop()
		card2 := deck2.Pop()

		// fmt.Printf("=====\n[Round %d] Play: [1] %d, [2] %d\n", rounds, card1, card2)

		if card1 > card2 {
			// fmt.Printf("Player 1 wins\n")
			deck1.Add(card1, card2)
		} else {
			// fmt.Printf("Player 2 wins\n")
			deck2.Add(card2, card1)
		}

		// fmt.Printf("Player1: %v\n", deck1)
		// fmt.Printf("Player2: %v\n", deck2)
	}

	fmt.Printf("Played %d rounds\n", rounds)
	fmt.Printf("Player1: %v\n", deck1)
	fmt.Printf("Player2: %v\n", deck2)

	var winner Deck
	if len(deck1) > 0 {
		winner = deck1
	} else {
		winner = deck2
	}

	score := 0
	for i, card := range winner {
		score += card * (len(winner) - i)
	}
	fmt.Printf("Final score: %d\n", score)
}

func readDecks(reader io.Reader) (Deck, Deck, error) {
	var deck1, deck2 Deck
	scanner := bufio.NewScanner(reader)

	var currentDeck *Deck
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Player 1:" {
			currentDeck = &deck1
		} else if line == "Player 2:" {
			currentDeck = &deck2
		} else if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, err
			}
			*currentDeck = append(*currentDeck, num)
		}
	}

	if scanner.Err() != nil {
		return nil, nil, scanner.Err()
	}

	return deck1, deck2, nil
}
