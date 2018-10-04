package main

import "fmt"

func main() {

	word := "hi"
	countWord(word)
}

func countWord(w string) int {
	// else clause must be on same line
	if len(w) != 2 {
		fmt.Println("Should equal 2")
	} else {
		fmt.Println("Adds up correctly")
	}
	return len(w)
}
