package main

func main() {
	cards := newDeck()
	dealCards, remainingCards := deal(cards, 5)

	dealCards.print()
	remainingCards.print()
}
