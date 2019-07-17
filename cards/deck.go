package main

import (
	"fmt"
)

// Create a new type of Deck, which is a slice of Strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades" , "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of"+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
