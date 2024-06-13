package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
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
	for _, link := range links {
		// Call the method as go routines
		go checkLink(link)
	}

	// Wait for all goroutines to finish
	time.Sleep(5 * time.Second)

}
