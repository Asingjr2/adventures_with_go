// Creating channel that only works in one direction
package main

import (
	"fmt"
)

func main() {
	// Basic Channel with buffer of 3
	c := make(chan int, 2)
	c <- 100
	c <- 200

	// SEND ONLY CHANNEL (receiving int)
	cSend := make(chan<- int, 2)
	go func() {
		cSend <- 100
	}()
	// fmt.Println(<- cSend)	// Creates error since we cannot receive!

	cReceive := make(<-chan string)
	// cReceive <- 100			// Creates error since we can only send!

	fmt.Println("______")
	fmt.Printf("%T\n", c)			// returns chan int
	fmt.Printf("%T\n", cSend)		// returns chan<- int
	fmt.Printf("%T\n", cReceive)	// returns <-chan string

	// Can assign directional to general (e.g. (chan<- int))(newChan)
}
