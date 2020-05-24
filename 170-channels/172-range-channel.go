package main

import "fmt"

func main() {
	c := gen()
	fmt.Printf("Your channel:%v type %T\n", c, c)
	receive(c)
	// close(c)

	fmt.Println("about to exit")
}

func receive(c <-chan int) {
	// Range over all values in a channel
	// NB that channels are always FIFO
	// Like a PIPE
	for val := range c {
		fmt.Println("Value:", val)
	}
}

func gen() <-chan int {
	// Make a channel
	// Push 0-99 into the channel
	// Return the channel, as a receive chan
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i // Push into channel
		}
		close(c)
	}()
	return c
}
