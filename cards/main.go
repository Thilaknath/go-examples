package main

import "fmt"

func main() {
	//var card string = "Ace of Spades" (very fist assignnent you use colon)
	cards := newDeck()
	fmt.Println(cards.toString())
}

