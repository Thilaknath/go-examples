package main

import (
	"fmt"
)

// Create a new type of Deck, which is a slice of Strings
// _ Used to replace a variable which is not used (In below example it was i and j in the for loop)
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades" , "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
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
