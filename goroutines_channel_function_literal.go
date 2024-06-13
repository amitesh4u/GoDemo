package main

import (
	"fmt"
	"net/http"
	"time"
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
		"https://dummy.abc",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
	/* Create Channel of Link since main function may exit before goroutines process completes
	Similar to Future in java */
	c := make(chan string)

	for _, link := range links {
		// Call the method as go routines
		go checkLinkWithLinkChannel(link, c)
	}
	/* Infinitely processing the status check by passing the link again
	to method recursively with a delay inside a function literal */
	for l := range c { // Set l with value of Channel c
		go func(link string) { // Set he parameter link with value of l
			time.Sleep(3 * time.Second) // Add a delay
			// recursively call the check link function with link received from channel
			go checkLinkWithLinkChannel(link, c)
		}(l) // Pass the value of l to function literal
	}

}
