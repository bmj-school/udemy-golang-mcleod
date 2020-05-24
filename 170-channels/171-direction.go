package main

import (
	"fmt"
)

func main() {
	// This is a SEND channel,
	// ints are PLACED into the channel
	cs := make(chan<- int)

	go func() {
		// Here, we want to receive, will not work!
		// cs <- 42
	}()

	// Instead, use a bidirectional

	c_bi := make(chan int)
	go func() {
		c_bi <- 42
	}()

	fmt.Println(<-c_bi)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
