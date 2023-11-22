package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic to generating english greeting
	return "Hi There !!!"
}

func (eb spanishBot) getGreeting() string {
	// very custom logic to generating english greeting
	return "Hola Hola !!!"
}
