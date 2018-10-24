// Addressing control flow which works like traditional langauge.
// Fallthrough may be unique to go.
package main

import "fmt"

func main() {
	// Basic loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("____________")
	// Basic loop 2
	stuff := []int{9, 20, 2342}
	for _, num := range stuff {
		fmt.Println(num)
	}

	fmt.Println("____________")
	// Basic loop with if-else statement.
	for i := 0; i < 11; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		} else if i%5 == 0 {
			fmt.Println("divisible by 5")
		} else {
			fmt.Println("not divisible")
		}
	}

	fmt.Println("____________")
	// Switch expression statement (missing expression looks for true case)
	switch {
	case false:
		fmt.Println("will not print")
	case true:
		fmt.Println("def prints")
		fallthrough // Should not use sine everything under will be considered true
	case (2 == 2):
		fmt.Println("will only print with fallthrough if previous case is true")
		fallthrough
	case false:
		fmt.Println("will also print when false with fallthrough")
	}

	fmt.Println("____________")
	// Swtich should have default
	switch "a" {
	case "B":
		fmt.Println("will not print")
	default:
		fmt.Println("prints since captial B is not lowercase a")
	}

}
