package main

func main() {
	cards := newDeck()

	cards.print()

	hand, remainingCards := cards.deal(5)

	hand.print()
	remainingCards.print()
}
