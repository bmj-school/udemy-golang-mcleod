package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	counter = 0
	fmt.Println("Counter:", counter)

	goroutines := 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for gr := 0; gr < goroutines; gr++ {
		// fmt.Println(gr)
		go func() {
			var thisctr = counter
			// fmt.Println(gr)
			runtime.Gosched()
			thisctr++
			counter = thisctr
			wg.Done()

		}()
	}
	fmt.Println("Counter:", counter)

	wg.Wait()
}
