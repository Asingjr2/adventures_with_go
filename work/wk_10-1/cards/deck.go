package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	// Creating two for loops to create deck
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// returns two deck objects of the divided original
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// changes information to string using std package
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// taking arguement to write new file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// using built in function from util to read file
	// converting returned byte slice => string => deck
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// if no errors than move forward with code
		// err code of 0 means nothing is wrong
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// swaps index with random num between 0 and len(deck)
	// Creating new seed using built in method that takes int64.  We use current time in seconds

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// newPos := rand.Intn(len(d) - 1)
		newPos := r.Intn(len(d) - 1) // function wtih custom r with new seed/source

		// format does not exist in python...need temp value to swap
		d[i], d[newPos] = d[newPos], d[i]
	}
}
