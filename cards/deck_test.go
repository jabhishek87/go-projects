package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen := 52
	if len(d) != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v ", expectedLen, len(d))
	}
}

func TestSaveNLoadDeck(t *testing.T) {

	expectedLen := 52
	// Step Create Deck
	d := newDeck()
	// Save to file
	filename := "_deckTesting.card"
	d.saveToFile(filename)

	// Load from file
	deckFromFile := newDeckFromFile(filename)

	// if type(d) != type(deckFromFile) {
	// 	t.Errorf("loaded file is not same as saved")
	// }

	if len(d) != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v ", expectedLen, len(d))
	}
	if len(d) != len(deckFromFile) {
		t.Errorf("not same")
	}
	os.Remove(filename)

}
