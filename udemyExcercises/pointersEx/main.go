// Ninja 7 (printing address of a variable)
// Other assignment is creating function that takes pointer and changes

package main

import (
	"fmt"
)

type person struct {
	First string
}

func main() {
	word := "word"
	fmt.Println(&word)

	newPerson := person{"Joe"}
	fmt.Println(newPerson)
	newPersonPointer := &newPerson
	changeMe(newPersonPointer)
	fmt.Println(newPerson)
}

func changeMe(p *person) {
	(*p).First = "Billy"
}
