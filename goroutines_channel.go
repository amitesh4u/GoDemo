package main

import (
	"fmt"
	"net/http"
)

func checkLinkWithStringChannel(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "might be down!")
		c <- link + " might be down!"
		return
	}

	//fmt.Println(link, "is up!")
	c <- link + " is up!"
}

func testGoRoutinesWithChannel() {
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
		go checkLinkWithStringChannel(link, c)
	}
	// Waiting and processing return of all channel results
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}
