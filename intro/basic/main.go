// Covering basic concepts and syntax
package main

import "fmt"

func main() {
	// base variable assignment
	var card1 string = "Ace of Spades"
	card2 := "Ace of Spades"
	card2 = "Three" // can change once iniitialized

	// creating a slice...which is an array that can expand...same types
	cards := []string{
		newCard(), newCard()}
	fmt.Println(cards)
	// adding to a slice
	cards = append(cards, "NEW ADDITION")
	fmt.Println(cards)

	// iteration over a slice with print index, and value
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(card1)
	fmt.Println(card2)
	newCard()
}

func newCard() string {
	// base function declaration with return type after like C#
	return "New Card Jack"
}
