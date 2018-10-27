// Map is equivalent to dictionaries
package main

import "fmt"

func main() {
	// creation of map requires static keys to be same and values to be same type
	// Key and value can be different types from each otyerh
	// commas are required after every line
	colors := map[string]string{
		"red":    "#ff000",
		"blue":   "$8022",
		"yellow": "#4314g",
	}
	fmt.Println(colors)

	// creating another map
	colors2 := make(map[string]string)

	// adding to map, not technically initializing so reg equals sign is ok
	colors2["brown"] = "#fzzzz"
	fmt.Println(colors2)

	// delete item using built in function with index coming second
	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)

}

func printMap(m map[string]string) {
	// function to work through items in map which is like dict
	for color, hex := range m {
		fmt.Println("Color is %v", color)
		fmt.Println("Hex is %v", hex)
	}

}
