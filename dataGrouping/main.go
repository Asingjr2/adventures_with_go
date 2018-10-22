//  Covering fundamental data groups in go language
package main

import "fmt"

func main() {
	// Array is fixed length list or same type
	var arr1 [3]int
	fmt.Println(arr1)
	arr1[0] = 1000
	fmt.Println(arr1)

	// Slice allows you to declare elements at creation
	slice1 := []int{89, 1000, -423, 3213}
	fmt.Println(slice1[1])
	// Can parts of slice
	sliceFragment := slice1[2:]
	fmt.Println(sliceFragment)

	// Looping through with range
	ppl := []string{"Joe", "Mary", "Chris", "Sunny", "Jenn"}
	for _, person := range ppl {
		fmt.Println(person)
	}

	// Variadic elements (.e.g. ...T means unlimited variables of T)
	sum(10, 11, 9)

	me := []string{"Samus"}	
	morePeople := append(ppl, "Ryan", "Jose")	
	evenMorePeople := append(me, morePeople...)		// Unfurling morePeople
	fmt.Println(morePeople)
	fmt.Println(evenMorePeople)
}

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	// Code loops through all entered nums and updates total
	fmt.Println(total)
}
