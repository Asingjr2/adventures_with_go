// Experimenting waitgroup which creates a default coutner that waits for a number
// of activities to finish.  wg.Add(1) gives total of 1 activity to finish
package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	arch = runtime.GOARCH
	os   = runtime.GOOS
)

var wg sync.WaitGroup // creating package level waitgroup

// Checking specific PC elements
func main() {
	fmt.Println("architecture", arch)
	fmt.Println("os", os)
	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("routines", runtime.NumGoroutine())

	// If run outright completes because main func code finishes first
	go play()

	wg.Add(1)
	foo()
	bar()

	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("routines", runtime.NumGoroutine())
	wg.Wait()
}

// Go routine set to run below which cannot complete before main finishes
// If max i dropped to 2 program could possible run before main finishes
func play() {
	for i := 0; i < 1000; i++ {
		fmt.Println("play", i)
	}
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo", i)
		wg.Done()
	}
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar", i)
	}
}
