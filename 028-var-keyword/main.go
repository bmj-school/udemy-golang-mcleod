package main

import "fmt"

var y = 43

func main() {
	// Short declaration
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	y = 44
	foo()
}

func foo() {
	fmt.Println(y)
}
