package main

import (
	"fmt"
)

// Create a new type of 'deck' which is a slice of
// strings
type deck []string

// not a receiver function because we won't have a
// deck yet to work with, since receiver funcs only
// work if you have an instance of that type created
// to do something with
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits { // _ means we know there's a variable there but we don't need to use it
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
