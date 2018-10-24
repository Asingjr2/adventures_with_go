// Fanin design pattern.
// Fans information from multiple channels into one
// Below code takes a even channel and odd channel and send both to fanin c
package main

import (
	"fmt"
	"sync"
)

func main() {
	// making basic bidirectional channels
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd) // distinct routine to send information into channel

	go receive(even, odd, fanin) // distinct routine to receive information

	for v := range fanin { // range loop to print information in channel and close when finished
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan int) { // function considers channel send channel only
	evenCount := 0
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // instruction to send information into channle thats even
			even <- i
			evenCount++
			fmt.Println("added to even", evenCount)
		} else { // instruction to send information into channle thats odd
			odd <- i
		}
	}
	close(even) // close channel when everything is done being sent
	close(odd)
}

// receive channel
// Even and odd are send only channel (we changed them previously)

// Fanin in a receive only channel
func receive(even, odd, fanin chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	// Function range over whats in eve and sends to fanin
	go func() {
		for v := range even {
			// fanin is considered receive only channel
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			// fanin considered receive only channel
			fanin <- v
		}
		wg.Done()
	}()

	// Waiting until all information is received in fanin
	wg.Wait()
	close(fanin)
}
