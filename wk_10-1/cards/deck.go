package main

import "fmt"

// declaring new data type
type deck []string

func (d deck) print() {
	// function contains a receiver...d deck... contains a type (e.g. deck) with name of "d"
	// best practice to name variable in one or two letters representing type
	// new functions name is print()
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	// function that creates and returns a new deck
	cards := deck{} // creating empty deck before appending
	suits := []string{"Spade", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	// Creating two for loops to create deck
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}
