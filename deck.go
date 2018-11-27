package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
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

// receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver function
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// receiver function
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(randomInt64Seed()) // uses current time to create source with a diff number with each call
	seededRand := rand.New(source)              // creates new rand instance so we can pass in a diff seed value and get better randomization

	for i := range d {
		newPosition := seededRand.Intn(len(d) - 1)  // get random number from 0 to length of slice
		d[i], d[newPosition] = d[newPosition], d[i] // swap cards
	}
}
