package main

import (
	"fmt"
)

func main() {
	testDeckFlows()

	testStructPointerFlows()

	testInterface()

	testHttpCall()

	testGoRoutines()
}

func testGoRoutines() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://dummy.abc",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
	// Create Channel of String since main function may exit before goroutines process completes
	// Similar to Future in java
	c := make(chan string)

	for _, link := range links {
		// Call the method as go routines
		go checkLink(link, c)
	}
	// Waiting and processing return of all channel results
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func testHttpCall() {
	makeGetCall()
}

func testInterface() {
	e := engBot{}
	printGreeting(e)

	s := spanishBot{}
	printGreeting(s)
}

func testStructPointerFlows() {
	person := Person{
		firstname: "fName",
		lastName:  "lName",
		ContactInfo: ContactInfo{
			mobileNo: "9876543210",
			email:    "test@test.com",
		},
	}

	fmt.Print(person) // Without Key
	//{fName lName {9876543210 test@test.com}}

	fmt.Printf("\n%+v", person) // With Key: Value
	//{firstname:fName lastName:lName ContactInfo:{mobileNo:9876543210 email:test@test.com}}

	personPointer := &person // Passing the memory address
	personPointer.updateFirstName("updatedFName")
	fmt.Printf("\n%+v", person)

	//OR in shortcut. Go will convert it to memory address if the receiver is a pointer
	person.updateFirstName("updatedAgainFName")
	fmt.Printf("\n%+v", person)
}

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
