package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Creates a new Type 'Deck' which is a slice of Strings
type deck []string

// (d deck) is receiver. Any var of type Deck get access to method Print
func (d deck) print() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~")
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
	var shuffleIndices = make(map[int]bool)
	var shuffledDeck = deck{}
	var cardLen = len(d)
	for len(shuffleIndices) < cardLen {
		index := randomGenerator().Intn(cardLen)
		if !shuffleIndices[index] {
			shuffledDeck = append(shuffledDeck, d[index])
			shuffleIndices[index] = true
		}
	}
	return shuffledDeck
}

// go lang uses same seed everytime so changing it to use current time
func randomGenerator() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano()) //Seed must be of type int64
	return rand.New(source)
}
