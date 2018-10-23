// Experimenting waitgroup which creates a default coutner that waits for a number
// of activities to finish.  wg.Add(1) gives total of 1 activity to finish
// If waitGroup not used in below code, Foo() will not ever complete since the program will close.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	// Checking specific PC elements
	arch = runtime.GOARCH
	os   = runtime.GOOS
)

var wg sync.WaitGroup // creating package level waitgroup

func main() {
	fmt.Println("architecture", arch)
	fmt.Println("os", os)
	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("routines", runtime.NumGoroutine())

	// If run outright completes because main func code finishes first
	wg.Add(1)
	go foo() // would not complete without telling program to wait
	bar()

	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("routines", runtime.NumGoroutine())
	wg.Wait()
}

// Go routine set to run below which cannot complete before main finishes

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
