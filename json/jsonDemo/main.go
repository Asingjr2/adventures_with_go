// Sample website call and parsing of json data.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ====================================
// Creating struct with "dummy" attribute to similulate additional data being returned
// Error checking in json only occurs against type of attribute listed
// Other returned JSON elements are ignored if not defined in struct
// ====================================
type serverTime struct {
	ServerTime111 int `json:"serverTime"` // If method type not matching json false value returned
	random        string
	random2       string
}

func main() {
	url := "https://api.binance.com/api/v1/time"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("didn't work", err)
	}

	defer resp.Body.Close()

	// Body is returned json object from call
	body, err := ioutil.ReadAll(resp.Body)

	// Error checking
	if err != nil {
		log.Fatal("body reference error", err)
	}

	var st serverTime
	// Unmarshall breaks down json object into individual parts
	err = json.Unmarshal(body, &st)

	// Error checking
	if err != nil {
		log.Fatal("response err", err)
	}

	fmt.Println(st.ServerTime111)
	fmt.Println(st.random)
	log.Printf("%+v", st)

}
