// Exploring tags and conversion
package main

import (
	"encoding/json"
	"fmt"

	dec "github.com/shopspring/decimal"
)

type animal struct {
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

type letters []string

type weight struct {
	Pounds dec.Decimal `json:"lbs"`
}

func main() {
	// *************************INSTANCIATING OBJECTS
	lion := animal{"Lion", 14}
	dog := animal{"Lab", 13}
	pack := []animal{lion, dog}

	// *************************MARSHAL INTO JSON

	// Marshal returns byte slice and err
	bsPack, err := json.Marshal(pack)
	if err != nil {
		fmt.Println(err)
	}
	stringPack := string(bsPack)

	fmt.Printf("%T\n", bsPack)     // type []uint8
	fmt.Printf("%T\n", stringPack) // type string
	fmt.Println(`bsPack`)          // returns string word bsPack
	fmt.Println(stringPack)        // returns json string

	// *************************UNMARSHALLING INTO OBJECTS
	var unmarshalPack []animal

	// Unmarshaling byteslice into array of objects
	err = json.Unmarshal(bsPack, &unmarshalPack)
	if err != nil {
		fmt.Println("error unmarshaling", err)
	}
	fmt.Println("printing updated umarshaled information\n", unmarshalPack)
	fmt.Printf("%T\n", unmarshalPack)

	// *************************TESTING TAGS
	fakeJSONWeight := `{"lbs":70.0321}`

	var convertedWeight weight
	err = json.Unmarshal([]byte(fakeJSONWeight), &convertedWeight)
	if err != nil {
		fmt.Println("error unmarshaling", err)
	}
	fmt.Println(convertedWeight)
	fmt.Printf("%T", convertedWeight)
}
