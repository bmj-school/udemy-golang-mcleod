package main

import "fmt"

var y = 43

func main() {
	// Short declaration
	x := 42
	fmt.Printf("%T", x)
	y = 44
	fmt.Printf("%T", y)
}