package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter uint64
	counter = 0

	fmt.Println("Counter:", counter)

	goroutines := 100

	var wg sync.WaitGroup

	for gr := 0; gr < goroutines; gr++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&counter, 1)
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
