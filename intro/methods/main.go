package main

import "fmt"

type thing struct {
	name, color string
}

func main() {
	first := thing{"box", "brown"}
	fmt.Println(first)
}
