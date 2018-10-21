// Major takeways from pointer concepts
// pointer goes to value with *value
// value goes to pointer with &value

package main

import "fmt"

type dog struct {
	name string
}

type cat struct {
	breed string
}

func main() {
	seven := 7
	fmt.Println(seven)
	sevenPointer := &seven // special is references the pointer of seven
	fmt.Println(sevenPointer)

	spike := dog{"Spike"}
	spike.newName("Rex") // function does not alter original value
	fmt.Println(spike.name)
	spikePointer := &spike
	spikePointer.newName2("Dot") // pointer being passed changes original
	fmt.Println(spike.name)

	spot := dog{"Spot"} // shortcut
	// Was able to base original value and pointer into function
	spot.newName2("spot changed to Dot")
	fmt.Println(spot.name)
}

func (d dog) newName(name string) {
	d.name = name
	fmt.Println(d.name)
}

// function takes different arguement
func (dogPointer *dog) newName2(name string) {
	// * in front of type means that we are using a pointer type
	(*dogPointer).name = name
}

