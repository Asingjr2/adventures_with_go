// Practice with marshal (turning json into values) and unmarshal (turning values into json)
package main

// Important resource: https://mholt.github.io/json-to-go/

import (
	"encoding/json"
	"fmt"
)

type (
	person struct {
		First string `json:"playValue"` // json tag overwrites returned value
		Last  string
		Age   int
	}
	// New type to check when unmarshaling
	adult struct {
		Second string `json:"playValue"`
		Last   string
		Age    int
	}
)

func main() {

	// Creating instances to marshal
	p1 := person{"James", "Bond", 32}
	p2 := person{"The", "Rock", 28}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	// Below line returns JSON object
	fmt.Println(string(bs))

	// Unmarshal parses json object
	// can also write as var adults []person
	adults := []person{}

	// Can write following line two ways.  Changing string json into bytes and then unmarshaling
	// err = json.Unmarshal(]byte(bs), &adults	)
	err = json.Unmarshal(
		[]byte(bs), &adults, // putting on new line requires comma
	)
	if err != nil {
		fmt.Println("unmarshal error", err)
	}

	fmt.Println("unmarshal worked", adults)
	for _, v := range adults {
		fmt.Println(v.First, v.Last, v.Age)
	} 
}
