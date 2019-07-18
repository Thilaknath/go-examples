package main

func main() {
	//var card string = "Ace of Spades" (very fist assignnent you use colon)
	cards := newDeck()
	cards.saveToFile("my_cards")
}
