package main

// Using alias instead of import name
// using dot notation to import into current package
// Using underscore will avoind "not used error", not preferrable.  
import (
	printing "fmt"
	// . "fmt"
	_ "log"
)

func main() {
	printing.Println("hey sir")
	// Println("hey sir")
}