package main

func main() {
	//var card string = "Ace of Spades" (very fist assignnent you use colon)
	cards := newDeck()

	deal(cards, 5)
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}

