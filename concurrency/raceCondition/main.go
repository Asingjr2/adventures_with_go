// Code below improves race condition which limited processes to two
// Original race condition: https://play.golang.org/p/FYGoflKQej
// Can check if race condition exists by running flag in command line
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("cpus", runtime.NumCPU())
	fmt.Println("routines", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		// Executing func literal or anonymous function

		go func() {
			v := counter
			// Putting process to sleep would allow another process to run
			// time.Sleep(time.Second)
			runtime.Gosched() // Yields processor so something else can run
			v++
			counter = v
			// Tells group this particular process is finished and removes itself from group
			wg.Done()
		}()
	}

	// Waits for processes added to group to finish whatever they are doing
	wg.Wait()
	fmt.Println("routines", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
