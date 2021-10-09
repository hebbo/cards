package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("size is wrong. got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("first card is wrong. got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("last card is wrong. got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	var deck_file_path string = "_decktesting"
	os.Remove(deck_file_path)

	d := newDeck()
	d.saveToFile(deck_file_path)

	loadedDeck := newDeckFromFile(deck_file_path)

	if len(loadedDeck) != 16 {
		t.Errorf("could not load deck successfully. got %v cards in deck", len(loadedDeck))
	}

	os.Remove(deck_file_path)
}
