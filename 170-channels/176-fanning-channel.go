package main

import "fmt"

func main() {
	c := make(chan int)
	numRoutines := 10

	for gr := 0; gr < numRoutines; gr++ {

		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
			close(c)
		}()
	}

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
