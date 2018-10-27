// Creating bird interface in golang
package main

import (
	"fmt"
)

type (
	cardinal struct{}
	blueJay  struct{}
	bird     interface {
		whatTypeAmI() 
	}
)

func (c cardinal) whatTypeAmI() {
	fmt.Println("I am a cardinal")
}

func (b blueJay) whatTypeAmI() {
	fmt.Println("I am a bluejay")
}

func myTypeIs(niceBird bird) {
	fmt.Println("I am a glorious BIRDDDD")
}

func main() {
	// Creating
	card := cardinal{}
	blue := blueJay{}

	card.whatTypeAmI()
	blue.whatTypeAmI()

	myTypeIs(card)
	myTypeIs(blue)
}// Creating bird interface in golang
package main

import (
	"fmt"
)

type (
	cardinal struct{}
	blueJay  struct{}
	bird     interface {
		whatTypeAmI() 
	}
)

// Creating method required to also be considered a bird
func (c cardinal) whatTypeAmI() {
	fmt.Println("I am a cardinal")
}

// Creating method required to also be considered a bird
func (b blueJay) whatTypeAmI() {
	fmt.Println("I am a bluejay")
}

// Creating a function that takes a type interface.  
func myTypeIs(niceBird bird) {
	fmt.Println("I am a glorious BIRDDDD")
}

func main() {
	// Creating instances of a cardinal and bluejay
	card := cardinal{}
	blue := blueJay{}

	// Implementing same method with overloading since signature is diff
	card.whatTypeAmI()
	blue.whatTypeAmI()

	// Implementing interface specific func.  Since both types have required methods
	// Running this will work with no problem
	myTypeIs(card)
	myTypeIs(blue)
}
