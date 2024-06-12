package main

func main() {
	// Create a new Deck
	cards := newDeck()
	cards.print()

	// Shuffle the deck
	cards = cards.shuffle()
	cards.print()

	// Deal a hand
	hand := cards.deal()
	hand.print()

	// Save the deck to a file
	cards.saveToFile()

	// Recreate the deck from the file
	newCards := newDeckFromFile()
	newCards.print()

}
