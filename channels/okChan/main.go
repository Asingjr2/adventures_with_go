// ,Ok idiom checks to see if channel is closed (almost like error checking)
package main

import "fmt"

func main() {
	// Creating channel moving int
	c := make(chan int)
	// Literal or anonymous function
	go func() {
		c <- 42
		// Below closes channel manually.  If missing and nothing in channel no errors will occur
		// close(c) 
	}()

	// Assigning channel value to an int and checking if closed
	v, ok := <-c

	// Printing value of variable and ok value (which is a bool!!!)
	fmt.Println(v, ok)
}
