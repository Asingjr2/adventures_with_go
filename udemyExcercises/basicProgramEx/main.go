package main

import "fmt"

const (
	// Typed and untyped const
	a = 7
	b int = 8
)

func main() {
	// Raw string literal
	rawString := `words
	that 
	extend`
	fmt.Print(rawString, "\n")
	// Operators
	num := 100
	a := (num == 101)
	b := (num <= 200)
	c := (num >= 200)
	d := (num != 101)
	e := (num > 50)
	f := (num < 50)
	fmt.Println(a,b,c,d,e,f)
}