// Interface determined by method set for type
// Can replace identical named methods for shared use
package main

import "fmt"

type (
	// Creating bare structs so we can just build methods for them
	englishBot struct{}

	spanishBot struct{}

	// Creating method that will be shared by types with similiar functions
	// Anything that has greeeting method is considered a bot
	bot interface {
		greeting() string
	}
)

func main() {
	// Creating instances of bot that have no methods/ attributes
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// Specific to englishBot
func (eb englishBot) greeting() string {
	return "Sup"
}

// Specific to spanishBot
func (sb spanishBot) greeting() string {
	return "Hola"
}

// Function to print greeting that will repalced by interface
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.greeting())
// }

// Function to print greeting replaced by interface
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.greeting())
// }

// Function takes a type bot.  Because english and spanish spot are also type bot.  they can be used.
func printGreeting(b bot) {
	fmt.Println(b.greeting())
}
