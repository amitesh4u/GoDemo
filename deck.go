package main

import "fmt"

// Creates a new Type 'Deck' which is a slice of Strings
type deck []string

// (d deck) is receiver. Any var of type Deck get access to method Print
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Returns a new Deck. No receiver required here
func newDeck() deck {
	cardSuites := []string{"Heart", "Spades", "Club", "Diamond"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}
	// blank identifier _, indicates that we want to ignore it.
	// Else it will create Unused variable error
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, suite+" of "+value)
		}
	}
	return cards
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
