package main

func main() {
	//var card string = "Ace of Spades" (very fist assignnent you use colon)
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
