package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hello my name is", p.first, p.last)
}

func main() {
	fmt.Println(simple())
	fmt.Println(multi_return())

	x := []int{1, 2, 3, 2, 2}
	fmt.Println(variadic(x...))
	fmt.Println(arraySum(append(x, 1, 1)))

	deferFunc()

	p1 := person{"Bill", "Twosons", 26}
	p1.speak()
}

func simple() int {
	return 42
}

func multi_return() (int, string) {
	return 32, "test"
}

func variadic(x ...int) int {
	fmt.Println(x)
	var mysum int
	for i := range x {
		mysum += x[i]
	}
	return mysum
}

func arraySum(x []int) int {
	var mysum int
	for i := range x {
		mysum += x[i]
	}
	return mysum
}

func deferFunc() {

	defer fmt.Println("This line is defered until the end!")
	fmt.Println("This is the second line in the code")
}
