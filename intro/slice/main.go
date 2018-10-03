package main

import "fmt"

func main() {
	ppl := []string{"Joe", "Mary", "Chris", "Sunny", "Jenn"}
	for _, person := range ppl {
		fmt.Println(person)
	}

	fmt.Println(ppl[0:2])
}
