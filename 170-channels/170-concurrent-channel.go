package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	// This will break, since it needs to run concurrently (goroutine!)
	// c <- 42

	go func() {
		c <- 42
	}()

	// Alternatively, the channel could be buffered
	// c := make(chan int, 1)

	fmt.Println(<-c)
}
