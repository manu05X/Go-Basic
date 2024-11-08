package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" { // Corrected assertion for the last card
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 160 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	// if loadedDeck[0] != "Ace of Spades" {
	// 	t.Errorf("Expected first card of Ace of Spades, but got %v", loadedDeck[0])
	// }
	// if loadedDeck[len(loadedDeck)-1] != "Ace of Spades" {
	// 	t.Errorf("Expected last card of Ace of Spades, but got %v", loadedDeck[len(loadedDeck)-1])
	// }

	os.Remove("_decktesting")
}
