package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	d := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	h := d[:handSize]
	n := d[handSize:]

	return h, n
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
