package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	counter = 0

	fmt.Println("Counter:", counter)

	goroutines := 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	var mu sync.Mutex

	for gr := 0; gr < goroutines; gr++ {
		go func() {
			mu.Lock()
			thisctr := counter
			// runtime.Gosched()
			thisctr++
			counter = thisctr
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
