package main

import "fmt"

type Cards struct {
	suite string
	val   string
	name  string
}

type Deck []Cards

func newDeck() Deck {
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}
	objDeck := Deck{}

	for _, suite := range cardSuites {
		for _, val := range cardValues {
			objDeck = append(objDeck, Cards{suite: suite, val: val, name: fmt.Sprintf("%s of %s", val, suite)})
		}
	}

	return objDeck
}

func (d Deck) print() {
	for idx, card := range d {
		fmt.Println(idx, card.name)
	}
}
