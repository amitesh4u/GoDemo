package main

func testDeckFlows() {
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
	remainingDeck.saveToFile("out/deck.txt")

	// Recreate the deck from the file
	newCards := newDeckFromFile("out/deck.txt")
	newCards.print()
}
