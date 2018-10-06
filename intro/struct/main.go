// Structs are similiar in shape to an object, but do not have self built in
// placing & before a variable name means give me memory address of variable
// placing * means give me the information at address so you can manipulate real info at varible, not copy
package main

import "fmt"

type person struct {
	// creating data type of struct
	firstName string
	lastName  string
	age       int
}

type contactInfo struct {
	zipcode int
}

type adult struct {
	// creating struct with another struct inside
	name string
	contactInfo
}

func main() {
	// creating 'instance' of person.  Note parenthesis like python
	// first way
	alex := person{"Alex", "Smith", 69}

	// second way
	chris := person{firstName: "Chris", lastName: "Johnson", age: 17}

	// third way; initializes nil value, "", 0, false, etc
	var mary person

	fmt.Println(alex, chris)
	fmt.Printf("%+v", mary)

	// creating struct with another struct inside (e.g. adult struct contains contact struct)
	jim := adult{name: "Jim", contactInfo: contactInfo{zipcode: 22041}}
	fmt.Printf("%+v", jim)

	// working with pointers and new syntax
	jimPointer := &jim
	fmt.Println(jimPointer)
	jimPointer.updateName("JOEY")
	jimPointer.print()

	// creating a slice of structs
	stuff := []struct {
		x, y int
	}{
		{5, 80},
		{799, 90},
	}
	fmt.Println(stuff)
}

func (a adult) print() {
	fmt.Printf("%+v", a)
}

func (pointerToPerson *adult) updateName(name string) {
	// updating information using pointer
	(*pointerToPerson).name = name
}
