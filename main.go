package main

func main() {
	// Create a new Deck
	cards := newDeck()
	cards.print()

	// Shuffle the deck
	cards = cards.shuffle()
	cards.print()

	// Deal a hand
	firstHand, remainingDeck := deal(cards, 3)
	firstHand.print()
	remainingDeck.print()

	// Deal another hand
	secondHand, remainingDeck := deal(remainingDeck, 3)
	secondHand.print()
	remainingDeck.print()

	// Save the deck to a file
	cards.saveToFile()

	// Recreate the deck from the file
	newCards := newDeckFromFile()
	newCards.print()

}
