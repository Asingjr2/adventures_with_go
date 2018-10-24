// Using select statement to decide which values from channel will be used
// Select works like switch on it removes values from channels
// Select pulls value ready to come off channel off
package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)
}

// Creating special func to receivce int from a chan
func receive(e, o, q <-chan int) {
	// Creating infinite loop.
	for {
		// Select case address every item that can come through allowing channels to not block
		select {
		// Case to receive information from e
		case v := <-e:
			fmt.Println("came from eve", v)
		case v := <-o:
			fmt.Println("came from eve", v)
		case v := <-q:
			fmt.Println("came from quit", v)
			return
		}
	}
}
func send(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
