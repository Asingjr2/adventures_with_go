// Checking that a channel is closed with ",ok"
package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		c <- "Mario"
		c <- "Luigi"
	}()

	v, ok := <-c
	fmt.Println(v, ok)

}
