// Works a bit like C# and other static langauges with type definition
package main

func main() {
	// creating deck with for loops
	cards := newDeck()
	cards.saveToFile("all_cards")

	// creating deck object from saved file with separated values
	cards2 := newDeckFromFile("all_cards")
	cards2.print()

	cards2.shuffle()
	cards2.print()

}

func newCard() string {
	return "Five"
}
