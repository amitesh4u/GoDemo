package main

import (
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		//fmt.Println(link, "might be down!")
		c <- link + " might be down!"
		return
	}

	//fmt.Println(link, "is up!")
	c <- link + " is up!"
}
