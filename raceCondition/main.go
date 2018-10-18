// Experimenting with race condition dealing with shared info in routines
package main

import (
	"fmt"
	"runtime"
)

const (
	arch = runtime.GOARCH
	os   = runtime.GOOS
)

// Checking specific PC elements
func main() {
	fmt.Println("architecture", arch)
	fmt.Println("os", os)
	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("routines", runtime.NumGoroutine())

	// Never completes because main func code finishes first
	go foo()
	bar()

	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("routines", runtime.NumGoroutine())
}

// Go routine set to run below which cannot complete before main finishes
// If max i dropped to 2 program could possible run before main finishes
func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("foo", i)
	}
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar", i)
	}
}
