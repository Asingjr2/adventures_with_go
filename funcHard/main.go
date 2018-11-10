package main

import (
	"fmt"
)

type (
	Genre struct {
		Cat string
	}

	Book struct {
		BookGenre Genre
		Name      string
	}

	Collection struct {
		TotalBooks int
	}
)

func main() {
	horror := Genre{"Horror"}
	scaryBook := Book{
		BookGenre: horror,
		Name:      "IT",
	}
	myCollection := Collection{TotalBooks: 100}

	fmt.Println(horror)                  // returns Horror
	fmt.Println(scaryBook.BookGenre.Cat) // returns Horror
	fmt.Println(myCollection)

	// myCollection.printHorror("i love reading")
	myCollection.printHorror("i love reading", &horror)

}

func (c Collection) printHorror(randomWords string, paramsGenre *Genre) {
	if paramsGenre != nil {
		fmt.Println("\n\nthere are params")
	}
	fmt.Println("empty")
}
