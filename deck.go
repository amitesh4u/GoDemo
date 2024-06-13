package main

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strings"
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
	cardSuites := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}
	// blank identifier _, indicates that we want to ignore it.
	// Else it will create Unused variable error
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
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

// Return N cards from Top and return left over the cards
// Passing Deck as parameter here for novelty and returning multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Save deck to a local file
func (d deck) saveToFile(fileName string) {
	fileData := []byte(d.toString())
	//fmt.Println(fileData)
	err := os.WriteFile(fileName, fileData, fs.FileMode(os.O_RDWR))
	if err != nil {
		defer os.Exit(1)
		log.Fatal(err)
	}
}

// Convert a List into a delimited String
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Fetch Deck from local file
func newDeckFromFile(fileName string) deck {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return deck{}
	}

	return strings.Split(string(fileData), ",")
}
