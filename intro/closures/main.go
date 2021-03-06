package main

import "fmt"

func intSeq() func() int {
	i := 0
	mem := &i
	fmt.Println(mem)
	return func() int {
		// access the x defined in intSeq scope
		i++
		return i
	}
}

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())

	bone := 9
	fmt.Println(bone)
	bone = 200
	fmt.Println(bone)
}
