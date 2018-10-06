package main

import "fmt"

func main() {
	// when defining constants cannot use initialize method
	const greeting = "hey"
	const b int = 90
	fmt.Println(greeting)
	fmt.Println(b)

}
