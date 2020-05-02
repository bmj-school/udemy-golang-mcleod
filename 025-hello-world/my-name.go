package main

import "fmt"

func main() {
	fmt.Println("Hello ")
	foo()
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("More")
}

func foo() {
	fmt.Println("In foo")
}
