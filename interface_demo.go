package main

import "fmt"

func testInterface() {
	e := engBot{}
	printGreeting(e)

	s := spanishBot{}
	printGreeting(s)
}

type bot interface {
	greeting() string
}

type engBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.greeting())
}

func (e engBot) greeting() string {
	return "hello"
}

func (s spanishBot) greeting() string {
	return "Hola"
}
