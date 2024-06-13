package main

import (
	"fmt"
	"net/http"
)

func checkLinkWithLinkChannel(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func testGoRoutinesWithChannelAndFuncLiteral() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
	// Create Channel of String since main function may exit before goroutines process completes
	// Similar to Future in java
	c := make(chan string)

	for _, link := range links {
		// Call the method as go routines
		go checkLinkWithLinkChannel(link, c)
	}
	// Infinitely processing the status check by passing the link again to method recursively
	for {
		go checkLinkWithLinkChannel(<-c, c)
	}

}
