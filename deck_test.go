package main

import (
	"math/rand"
	"os"
	"testing"
)

const deckSize = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != deckSize {
		t.Errorf("Expecting 52 cards but got only %d", len(d))
	}

	firstCard := "Ace of Hearts"
	if d[0] != firstCard {
		t.Errorf("Expecting first card to be %s", firstCard)
	}

	fourteenthCard := "Ace of Spades"
	if d[13] != fourteenthCard {
		t.Errorf("Expecting fourteenth card to be %s", fourteenthCard)
	}

	lastCard := "King of Diamonds"
	if d[deckSize-1] != lastCard {
		t.Errorf("Expecting last card to be %s", lastCard)
	}
}

func TestSaveAndReadFromFile(t *testing.T) {
	// Remove existing test file is present
	testFileName := "out/deck_test.txt"

	err := os.Remove(testFileName)
	if err != nil {
		// do nothing
	}

	deck := newDeck()
	deck.saveToFile(testFileName)

	newDeck := newDeckFromFile(testFileName)

	if len(deck) != len(newDeck) {
		t.Error("New deck is of different length.")
	}

	index := rand.Intn(deckSize)
	if deck[index] != newDeck[index] {
		t.Error("Restored cards are not in same order")
	}

	// Remove file after testing
	err = os.Remove(testFileName)
	if err != nil {
		// do nothing
	}
}
