package main

import "fmt"

type dog struct {
	breed string
}

func main() {
	// Value Types

	val1 := 1
	val2 := 2.05
	val3 := "three"
	val4 := true
	val5 := dog{"terrier"}
	simple()
	fmt.Println(val1, val2, val3, val4, val5)

	// running function to change
	valueType(val1, val2, val3, val4)
	val5.structChange()
}

func simple() string {
	return "simple"
}

func valueType(i int, f float64, s string, b bool) {
	i = 10
	s = "changed"
	f = 290.213
	b = false
	fmt.Println(i, f, s, b)
}

func (s dog) structChange() {
	s.breed = "lab"
	fmt.Println(s)
}
