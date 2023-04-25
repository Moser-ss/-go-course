package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	is := assert.New(t)
	d := newDeck()
	is.Len(d, 16)
	is.Equal("Ace of Spades", d[0])
	is.Equal("Four of Clubs", d[len(d)-1])

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	tf := "_decktesting"
	os.Remove(tf)
	deck := newDeck()
	deck.saveToFile(tf)

	loadedDeck := newDeckFromFile(tf)

	assert.Equal(t, deck, loadedDeck)
	os.Remove(tf)
}
