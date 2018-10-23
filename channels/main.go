// Exploring creation of chan and related buffers
package main

import (
	"fmt"
)

func main() {
	// Successful buffer
	// Creating channel receiving int with a buffer that can be held in channel without receiver
	c := make(chan int, 2)
	// Sending 42 into channel
	c <- 42
	c <- 420
	// Code would not run without buffer since chan would block
	fmt.Println(<-c)
	fmt.Println(<-c)

	c2 := make(chan int)

	go func() {
		c2 <- 52
	}()

	fmt.Println(<-c2)
}
