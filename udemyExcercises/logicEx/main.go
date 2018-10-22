package main

import "fmt"

func main() {
	// Which evaluate to true below:
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false || false)
	fmt.Println(true || false)
	fmt.Println(!false)

	// Switch condition
	favSport := "football"
	switch favSport {
	case "soccer":
		fmt.Println("soccer is best")
	case "tennis":
		fmt.Println("tennis is best")
	default:
		fmt.Println("football is best")
	}
}
