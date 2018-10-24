// Loop goes on forever until channel is closed. Range works through channel.
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 101)

	// send information through channel
	go func() {
		for i := 0; i < 25; i++ {
			c <- i
		}
		// If C is not closed the system is deadlocked
		close(c)
	}()

	// receiving information from channel.  will waiting until inforamtion is sent
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
