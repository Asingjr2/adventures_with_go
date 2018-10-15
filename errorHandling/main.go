// Addressing error handling methods with log, fatal, panic
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Defining error instance that can be referenced later
var ErrNot4 = errors.New("wrong number: Square root not equal to 4")

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err occured", err)
		log.Println("err occurred with date", err)
		// log.Fatalln(err)  will exit program and give error info
		// panic(err)
	}

	_, err2 := Sqrt4(5) //  returns error variable
	if err != nil {
		fmt.Println(err2) // returns nil
	}

	_, err3 := Sqrt4(16)
	if err != nil {
		fmt.Println(err3)
	}
}

func init() {
	// Creating log that will store errors
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func Sqrt4(f float64) (float64, error) {
	if f != 16 {
		return 0, ErrNot4
	}
	return 42, nil
}
