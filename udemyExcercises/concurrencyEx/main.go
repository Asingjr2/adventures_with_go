//  Using MUTEX or mutual exclusion lock.
// Program runs on my PC but does not in playground or others presumably
// Mutex locks a file for use by a process
package main

import (
	"fmt"
	// "sync"
)

// var wg sync.WaitGroup

func main() {
	// wg.Add(2)
	fmt.Println("starting other routines")
	
	go foo()
	go bar() 
	fmt.Println("done")
	// wg.Wait()
}

func foo(){
	count := 100
	for i := 0; i < count; i++ {
		fmt.Println("foo")
	}
	// wg.Done()
}

func bar(){
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("bar")
	}
	// wg.Done()
}