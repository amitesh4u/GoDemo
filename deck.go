package main

import "fmt"

// Creates a new Type 'Deck' which is a slice of Strings
type deck []string

// (d deck) is receiver. Any var of type Deck get access to method Print
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Returns a new Deck. No receiver required here
func newDeck() deck {
	return deck{}
}

// Shuffle the Deck
func (d deck) shuffle() deck {
	return deck{}
}

// Return a hand of cards
func (d deck) deal() deck {
	return deck{}
}

// Save deck to a local file
func (d deck) saveToFile() {

}

// Fetch Deck from local file
func newDeckFromFile() deck {
	return deck{}
}
