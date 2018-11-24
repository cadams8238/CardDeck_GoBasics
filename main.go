package main

import (
	"fmt"
)

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Five of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
