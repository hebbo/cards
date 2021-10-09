package main

import "testing"

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
