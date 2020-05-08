package main

import "fmt"

type MyType int
var x MyType

func main() {
	fmt.Printf("%T %i\n",x)
	x = 42
	fmt.Printf("%T %i\n",x)
}