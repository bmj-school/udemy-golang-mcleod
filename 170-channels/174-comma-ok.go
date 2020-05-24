package main

import (
	"fmt"
)

//The boolean variable ok returned by a receive operator indicates whether the
// received value was sent on the channel (true) or is a zero value returned
// because the channel is closed and empty (false).

// Therefore, we can detect if the channel is closed and empty, and handle it

func main() {
	c := make(chan int)
	go func() {
		c <- 32
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
