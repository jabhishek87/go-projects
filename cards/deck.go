package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type 'deck'
// which is slice of strings

type deck []string

// func to init new deck
func newDeck() deck {
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}
	cards := deck{}

	// Cretae deck with loop
	for _, suite := range cardSuites {
		for _, val := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", val, suite))
		}
	}
	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",,")

}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toByteSlice(), 0644)
}

func newDeckFromFile(filename string) deck {

	byte_, err := os.ReadFile(filename)
	if err != nil {
		// we can do to things
		// 1) log the error and quit the program
		// 2) log the error and return a call to new deck

		// we are going to implement option 1 for now
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(byte_), ",,")
	return deck(s)
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
