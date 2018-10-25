// Get basic channel running
// Experimenting with channel
package main

import "fmt"

func main() {

	// Making code run with buffer and o
	//
	c1 := make(chan int)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)

	go func() {
		c1 <- 100
		c2 <- 200
		c2 <- 2500
		c3 <- 300
		c3 <- 400
		// Cannot use range unless we close a chanel because range needs to know that a channl is not receiving anything else.
		close(c3)
	}()

	// Code will block until something is ready to come off
	fmt.Println(<-c1)
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	rangeChannel(c3)
}

func rangeChannel(sampleChannel <-chan int) {
	for v := range sampleChannel {
		fmt.Println(v)
	}
}
