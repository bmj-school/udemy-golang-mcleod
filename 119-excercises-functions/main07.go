package main

import "fmt"

func main() {
	func(){
		fmt.Println("He")
	}()
	this := func(){
		fmt.Println("Two")
	}

	this()
	fmt.Printf("%T, \n", this)

	fmt.Println(myCallback())
	thisFunc := myCallback
	fmt.Printf("%T, %v\n", thisFunc(), thisFunc())
}

func myCallback() func() int {
	return func() int {
		return 32
	}
}

