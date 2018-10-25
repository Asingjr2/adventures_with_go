// Interface determined by method set for type
package main

import "fmt"

type (
	// Creating bare structs so we can just build methods for them
	englishBot struct{}

	spanishBot struct {}
)

func main() {
	
}

func (e englishBot) greeting() string {
	return "Sup"
}

func (s spanishBot) greeting() string {
	return "Hola"
}