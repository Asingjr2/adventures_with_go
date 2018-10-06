package main

import "fmt"

func main() {
	// basic for
	sum := 2
	for i := 0; i < sum; i++ {
		fmt.Println(sum)
	}

	// shorthand
	count := 5
	for count < 8 {
		fmt.Println("Samus")
		count += 1
	}

	// for as while
	total := 2
	for total < 4 {
		fmt.Println("greatness")
		total += 1
	}

	// infinite loop
	// for {
	// do something
	// }

	// switch statement works like if statement
	name := "Art"
	switch len(name) {
	case 4:
		fmt.Println("picked 4")
	case 3:
		fmt.Println("name length is 3")
	default:
		fmt.Println("picked nothing")
	}

	// switch statement with no condition like long if statements
	switch {
	case name == "joe":
		fmt.Println("name is joe")
	case name == "mary":
		fmt.Println("name is mary")
	default:
		fmt.Println("name not joe or mary")
	}

	// defer keyword delays execution
	fmt.Println("first")
	defer fmt.Println("defer")
	fmt.Println("third")

}
