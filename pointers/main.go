// Covering what pointers are and syntax
// Use based on not moving large objects unless needed
package main

import (
	"fmt"
)

type special struct {
	word string
}

func main() {
	samus := "Samus"
	samus2 := &samus // &denotes memory address

	// Important distinctions being made below
	fmt.Println("Samus type \t, samus2 type, samus2 memory value")
	fmt.Printf("%T\t\t", samus)
	fmt.Printf("%T\t\t", samus2)
	fmt.Printf(*samus2)

	// *int is a data type, &variable is a pointer to an address, *variable returns value of information in memory

	num := 14
	fmt.Println("\n", num+10)

	// type instance
	ss := special{"Bowser"}
	fmt.Println(ss.word)
	ss.wordChange()
	ss.word = "Mario"
	fmt.Println(ss.word)

	ssPointer := &ss
	ssPointer.wordChange2()

	ssPointer.wordChange3()
	ss.word = "Mario"
	fmt.Println(ss.word)
	ss.wordChange3()

}

func (s special) wordChange() {
	s.word = "Samus"
	fmt.Println(s.word)
}

func (s special) wordChange2() {
	s.word = "Samus22222"
	fmt.Println(s.word)
}

func (s *special) wordChange3() {
	s.word = "Samus22222"
	fmt.Println(s.word)
}
