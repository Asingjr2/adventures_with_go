package main

import "fmt"

func main() {
	cards := deck{newCard(), newCard()}
	cards = append(cards, "SOME NEW CARD")
	fmt.Println(cards)

	cards.print()
}

func newCard() string {
	return "Five"
}

// declaring new data type
type deck []string

// declaring new function for type with a receiver of type 'instance'
func (d deck) print() {
	// function contains a receiver...d deck... contains a type (e.g. deck) with name of "d"
	// best practice to name variable in one or two letters representing type
	// new functions name is print()
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// declaring new function for type
func newDeck() deck {
	// function that creates and returns a new deck
	cards := deck{} // creating empty deck before appending
	suits := []string{"Spade", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	// Creating two for loops to create deck
	for i, suit := range suits {
		for j, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}
