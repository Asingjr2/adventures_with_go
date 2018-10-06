package main

import "fmt"

func main() {
	seven := 7
	fmt.Println(seven)
	special := &seven // special is references the pointer of seven
	fmt.Println(special)

}
