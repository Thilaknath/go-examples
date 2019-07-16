package main

import (
	"fmt"
)

func main() {
	//var card string = "Ace of Spades" (very fist assignnent you use colon)
	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i,card := range cards{
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
