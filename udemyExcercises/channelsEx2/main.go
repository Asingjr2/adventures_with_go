// Using a select statement with a channel
// Select works like switch but is required to adderss all items in a channel
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)
	fmt.Printf("chan typed change to return only when returned in function, \n")
	fmt.Printf("%T\n", c)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}

	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	fmt.Printf("%T\n", c)
	go func() {
		for i := 0; i < 4; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	fmt.Printf("getting basic chan type below, \n")
	fmt.Printf("%T\n", c)
	return c
}
