// Creating base file to use with testing.
// Testing requires manually entered values.

package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2,3))
	fmt.Println("5 + 10 = ", mySum(5,10))
	fmt.Println("7 + 300 = ", mySum(7,300))
}

func mySum(xi ...int) int {
	sum := 0
	for _,v := range xi {
		sum += v
	}
	return sum
}
