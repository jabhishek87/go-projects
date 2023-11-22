package main

import "fmt"

func main() {
	fmt.Println("Cards Struct")
	deck := newDeck()
	deck.print()
}
