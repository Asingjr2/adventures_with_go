// Ninja 8
// Marshal data
// Unmarshall data
// Send encoded data to stdOut by creating new encoding

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	first string
	age   int
}

type thing struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	u2 := user{
		first: "Moneypenny",
		age:   27,
	}

	u3 := user{
		first: "M",
		age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// Marshal
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("there was an err", err)
	}
	fmt.Println(bs) // prints byte slice

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var morePeople []thing
	// Unmarshal takes a byte slice and a POINTER to a place to put information
	err = json.Unmarshal([]byte(s), &morePeople)
	if err != nil {
		fmt.Println("\n\nerror unmarshalling", err)
	}
	fmt.Println(morePeople)

	// poniter to encoder
	err = json.NewEncoder(os.Stdout).Encode(users)

}
