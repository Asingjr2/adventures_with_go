// Works a bit like C# and other static langauges with type definition
package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards)

	cards.print()

}

func newCard() string {
	return "Five"
}
