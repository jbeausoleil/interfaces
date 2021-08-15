package main

import "fmt"

type bot interface { // any type in the program that calls getGreeting and returns a string is of type bot
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{} // also of type bot
	sb := spanishBot{} // also of type bot

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting (b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}