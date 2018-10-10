//  Experimenting with channels which much share same data type
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// creating channel which needs explicit type
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// calling function with returned link coming out of channel
	// creating function literal (e.g. lambda) with a callback
	for l := range c {
		go func(link string) {
			time.Sleep(1*time.Second)
			checkLink(link, c)
		}(l)  // l is being based into function literal
	}

	// above can be written as below.  variable is whatever comes out of channel
	for {
		go checkLink(<-c, c)
	}
}

// chan data type must be state explicitly
func checkLink(link string, c chan string) {
	// checks status of links in order with server check
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down %v", err)
		c <- link	// link being sent into channel
		return
	}
	fmt.Println(link, "link is good")
	c <- link
}
