package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 101)

	// send information through channel
	go foo(c)

	// receiving information from channel.  will waiting until inforamtion is sent
	bar(c)

	fmt.Println("about to exit")
}

// Sending information into channel.  If we create large group the program will close
// before all information is sent.  Using just 10 sends
func foo(c chan<- int) {
	// high := 100
	// for i := 0; i < high; i++ {
	// 	c <- i
	// }
	c <- 10
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
