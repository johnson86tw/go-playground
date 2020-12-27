package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{ lang string }
type spanishBot struct{ lang string }

type aStr string

func main() {
	e := englishBot{lang: "english"}
	s := spanishBot{lang: "spanish"}

	var a aStr = "a string"

	printGreeting(e)
	printGreeting(s)

	printGreeting(a)
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (englishBot) getBye() string {
	return "goodbye"
}

func (s spanishBot) getGreeting() string {
	return s.lang + "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (a aStr) getGreeting() string {
	return "I am just a string"
}
