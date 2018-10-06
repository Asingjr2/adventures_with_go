// exploring value versus reference type through updates to variables
// map are reference types so changes in functions alter original info
// structs are reference types so changes in functions update copy or pointer
package main

import (
	"fmt"
)

type person struct {
	name, food string
}

func main() {
	day := "mon"
	day2 := &day      // creating pointer to day
	fmt.Println(day2) // printing out mem location for pointer

	rin := map[string]string{
		"name": "rin",
		"food": "egg",
	}
	rinrin := &rin
	fmt.Println(rinrin)
	fmt.Println(rin) // does not print out mem address. Instead prints out &map....

	tito := person{"tito", "banana"}
	titotito := &tito
	fmt.Println(titotito) // does not print out address.  Instead prints out &{tito values}
	fmt.Println(tito)

	mapChange(rin)
	fmt.Println("%v", rin)
	structChange(tito)
	fmt.Println("%v", tito)
}

func mapChange(m map[string]string) {
	// will update real/ original data
	mm := &m
	fmt.Println(mm)
	m["name"] = "NEW NAME MAP"
	fmt.Println(m)
}

func structChange(s person) {
	// will alter information in function or copy
	ss := &s
	fmt.Println(ss)
	s.name = "New Name STRUCT"
	fmt.Println(s)
}
