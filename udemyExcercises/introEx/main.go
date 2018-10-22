package main

import "fmt"

var a int
var b string
var c bool

var d int = 42
var e = "James Bond"

func main() {
	// below prints "zero" values for type, 0, blank, and false
	fmt.Print(a)
	fmt.Println(b)
	fmt.Println(c)

	// Formatted lines
	sol := fmt.Sprintf("%v\t%v\t", d, e)
	fmt.Print(sol)
}