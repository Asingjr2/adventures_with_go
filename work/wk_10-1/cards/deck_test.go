// Package requires testing module and method arguements
// t.Errorf is used to annotate expected value or notes
package main

import (
	"testing"
	"os"
	)
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// Way to display types of error
		t.Errorf("Expected deck length of 20 but got " + string(len(d)))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Wrong card, %v", d[len(d)-1])
	}
}

func TestNewDeckFromFileAndNewDeckFromFile(t *testing.T) {
	// testing through creation and deletion of files with io package
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, got %v, ", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}